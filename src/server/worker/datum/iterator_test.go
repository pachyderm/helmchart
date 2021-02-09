package datum

import (
	"fmt"
	"strings"
	"testing"

	"github.com/pachyderm/pachyderm/v2/src/client"
	"github.com/pachyderm/pachyderm/v2/src/internal/dbutil"
	"github.com/pachyderm/pachyderm/v2/src/internal/require"
	"github.com/pachyderm/pachyderm/v2/src/internal/testpachd"
	tu "github.com/pachyderm/pachyderm/v2/src/internal/testutil"
)

func TestIterators(t *testing.T) {
	t.Parallel()
	db := dbutil.NewTestDB(t)
	require.NoError(t, testpachd.WithRealEnv(db, func(env *testpachd.RealEnv) error {
		c := env.PachClient
		dataRepo := tu.UniqueString(t.Name() + "_data")
		require.NoError(t, c.CreateRepo(dataRepo))
		// Put files in structured in a way so that there are many ways to glob it.
		commit, err := c.StartCommit(dataRepo, "master")
		require.NoError(t, err)
		for i := 0; i < 50; i++ {
			require.NoError(t, c.PutFile(dataRepo, commit.ID, fmt.Sprintf("/foo%v", i), strings.NewReader("input")))
		}
		require.NoError(t, c.FinishCommit(dataRepo, commit.ID))
		// Zero datums.
		in0 := client.NewPFSInput(dataRepo, "!(**)")
		in0.Pfs.Commit = commit.ID
		t.Run("ZeroDatums", func(t *testing.T) {
			pfs0, err := NewIterator(c, in0)
			require.NoError(t, err)
			validateDI(t, pfs0)
		})
		// Basic PFS inputs
		in1 := client.NewPFSInput(dataRepo, "/foo?1")
		in1.Pfs.Commit = commit.ID
		in2 := client.NewPFSInput(dataRepo, "/foo*2")
		in2.Pfs.Commit = commit.ID
		t.Run("Basic", func(t *testing.T) {
			pfs1, err := NewIterator(c, in1)
			require.NoError(t, err)
			pfs2, err := NewIterator(c, in2)
			require.NoError(t, err)
			validateDI(t, pfs1, "/foo11", "/foo21", "/foo31", "/foo41")
			validateDI(t, pfs2, "/foo12", "/foo2", "/foo22", "/foo32", "/foo42")
		})
		// Union input.
		in3 := client.NewUnionInput(in1, in2)
		t.Run("Union", func(t *testing.T) {
			union1, err := NewIterator(c, in3)
			require.NoError(t, err)
			validateDI(t, union1, "/foo11", "/foo12", "/foo2", "/foo21", "/foo22", "/foo31", "/foo32", "/foo41", "/foo42")
		})
		// Cross input.
		in4 := client.NewCrossInput(in1, in2)
		t.Run("Cross", func(t *testing.T) {
			cross1, err := NewIterator(c, in4)
			require.NoError(t, err)
			validateDI(t, cross1,
				"/foo11/foo12", "/foo11/foo2", "/foo11/foo22", "/foo11/foo32", "/foo11/foo42",
				"/foo21/foo12", "/foo21/foo2", "/foo21/foo22", "/foo21/foo32", "/foo21/foo42",
				"/foo31/foo12", "/foo31/foo2", "/foo31/foo22", "/foo31/foo32", "/foo31/foo42",
				"/foo41/foo12", "/foo41/foo2", "/foo41/foo22", "/foo41/foo32", "/foo41/foo42",
			)
		})
		// Empty cross.
		in6 := client.NewCrossInput(in3, in0, in2, in4)
		t.Run("EmptyCross", func(t *testing.T) {
			cross3, err := NewIterator(c, in6)
			require.NoError(t, err)
			validateDI(t, cross3)
		})
		// Nested empty cross.
		in7 := client.NewCrossInput(in6, in1)
		t.Run("NestedEmptyCross", func(t *testing.T) {
			cross4, err := NewIterator(c, in7)
			require.NoError(t, err)
			validateDI(t, cross4)
		})
		// in[8-9] are elements of in10, which is a join input
		in8 := client.NewPFSInputOpts("", dataRepo, "", "/foo(?)(?)", "$1$2", "", false, false, nil)
		in8.Pfs.Commit = commit.ID
		in9 := client.NewPFSInputOpts("", dataRepo, "", "/foo(?)(?)", "$2$1", "", false, false, nil)
		in9.Pfs.Commit = commit.ID
		in10 := client.NewJoinInput(in8, in9)
		t.Run("Join", func(t *testing.T) {
			join1, err := NewIterator(c, in10)
			require.NoError(t, err)
			validateDI(t, join1,
				"/foo11/foo11",
				"/foo12/foo21",
				"/foo13/foo31",
				"/foo14/foo41",
				"/foo21/foo12",
				"/foo22/foo22",
				"/foo23/foo32",
				"/foo24/foo42",
				"/foo31/foo13",
				"/foo32/foo23",
				"/foo33/foo33",
				"/foo34/foo43",
				"/foo41/foo14",
				"/foo42/foo24",
				"/foo43/foo34",
				"/foo44/foo44")
		})

		in11 := client.NewPFSInputOpts("", dataRepo, "", "/foo1(?)", "$1", "", true, false, nil)
		in11.Pfs.Commit = commit.ID
		in12 := client.NewPFSInputOpts("", dataRepo, "", "/foo(?)1", "$1", "", true, false, nil)
		in12.Pfs.Commit = commit.ID
		in13 := client.NewJoinInput(in11, in12)
		t.Run("OuterJoin", func(t *testing.T) {
			join1, err := NewIterator(c, in13)
			require.NoError(t, err)
			validateDI(t, join1,
				"/foo10",
				"/foo11/foo11",
				"/foo12/foo21",
				"/foo13/foo31",
				"/foo14/foo41",
				"/foo15",
				"/foo16",
				"/foo17",
				"/foo18",
				"/foo19")
		})

		in14 := client.NewPFSInputOpts("", dataRepo, "", "/foo(?)(?)", "", "$1", false, false, nil)
		in14.Pfs.Commit = commit.ID
		in15 := client.NewGroupInput(in14)
		t.Run("GroupSingle", func(t *testing.T) {
			group1, err := NewIterator(c, in15)
			require.NoError(t, err)
			validateDI(t, group1,
				"/foo10/foo11/foo12/foo13/foo14/foo15/foo16/foo17/foo18/foo19",
				"/foo20/foo21/foo22/foo23/foo24/foo25/foo26/foo27/foo28/foo29",
				"/foo30/foo31/foo32/foo33/foo34/foo35/foo36/foo37/foo38/foo39",
				"/foo40/foo41/foo42/foo43/foo44/foo45/foo46/foo47/foo48/foo49")
		})

		in16 := client.NewPFSInputOpts("", dataRepo, "", "/foo(?)(?)", "", "$1", false, false, nil)
		in16.Pfs.Commit = commit.ID
		in17 := client.NewPFSInputOpts("", dataRepo, "", "/foo(?)(?)", "", "$2", false, false, nil)
		in17.Pfs.Commit = commit.ID
		in18 := client.NewGroupInput(in16, in17)
		t.Run("GroupDoubles", func(t *testing.T) {
			group2, err := NewIterator(c, in18)
			require.NoError(t, err)
			validateDI(t, group2,
				"/foo10/foo20/foo30/foo40",
				"/foo10/foo11/foo12/foo13/foo14/foo15/foo16/foo17/foo18/foo19/foo11/foo21/foo31/foo41",
				"/foo20/foo21/foo22/foo23/foo24/foo25/foo26/foo27/foo28/foo29/foo12/foo22/foo32/foo42",
				"/foo30/foo31/foo32/foo33/foo34/foo35/foo36/foo37/foo38/foo39/foo13/foo23/foo33/foo43",
				"/foo40/foo41/foo42/foo43/foo44/foo45/foo46/foo47/foo48/foo49/foo14/foo24/foo34/foo44",
				"/foo15/foo25/foo35/foo45",
				"/foo16/foo26/foo36/foo46",
				"/foo17/foo27/foo37/foo47",
				"/foo18/foo28/foo38/foo48",
				"/foo19/foo29/foo39/foo49")
		})

		in19 := client.NewPFSInputOpts("", dataRepo, "", "/foo(?)(?)", "$1$2", "$1", false, false, nil)
		in19.Pfs.Commit = commit.ID
		in20 := client.NewPFSInputOpts("", dataRepo, "", "/foo(?)(?)", "$2$1", "$2", false, false, nil)
		in20.Pfs.Commit = commit.ID

		in21 := client.NewJoinInput(in19, in20)
		in22 := client.NewGroupInput(in21)
		t.Run("GroupJoin", func(t *testing.T) {
			groupJoin1, err := NewIterator(c, in22)
			require.NoError(t, err)
			validateDI(t, groupJoin1,
				"/foo11/foo11/foo12/foo21/foo13/foo31/foo14/foo41",
				"/foo21/foo12/foo22/foo22/foo23/foo32/foo24/foo42",
				"/foo31/foo13/foo32/foo23/foo33/foo33/foo34/foo43",
				"/foo41/foo14/foo42/foo24/foo43/foo34/foo44/foo44")
		})

		in23 := client.NewPFSInputOpts("", dataRepo, "", "/foo(?)(?)", "", "", false, false, nil)
		in23.Pfs.Commit = commit.ID
		in24 := client.NewPFSInputOpts("", dataRepo, "", "/foo(?)(?)", "", "$2", false, false, nil)
		in24.Pfs.Commit = commit.ID

		in25 := client.NewGroupInput(in24)
		in26 := client.NewUnionInput(in23, in25)

		t.Run("UnionGroup", func(t *testing.T) {
			unionGroup1, err := NewIterator(c, in26)
			require.NoError(t, err)
			validateDI(t, unionGroup1,
				"/foo10/foo20/foo30/foo40",
				"/foo11/foo21/foo31/foo41",
				"/foo12/foo22/foo32/foo42",
				"/foo13/foo23/foo33/foo43",
				"/foo14/foo24/foo34/foo44",
				"/foo15/foo25/foo35/foo45",
				"/foo16/foo26/foo36/foo46",
				"/foo17/foo27/foo37/foo47",
				"/foo18/foo28/foo38/foo48",
				"/foo19/foo29/foo39/foo49",
				"/foo10",
				"/foo11",
				"/foo12",
				"/foo13",
				"/foo14",
				"/foo15",
				"/foo16",
				"/foo17",
				"/foo18",
				"/foo19",
				"/foo20",
				"/foo21",
				"/foo22",
				"/foo23",
				"/foo24",
				"/foo25",
				"/foo26",
				"/foo27",
				"/foo28",
				"/foo29",
				"/foo30",
				"/foo31",
				"/foo32",
				"/foo33",
				"/foo34",
				"/foo35",
				"/foo36",
				"/foo37",
				"/foo38",
				"/foo39",
				"/foo40",
				"/foo41",
				"/foo42",
				"/foo43",
				"/foo44",
				"/foo45",
				"/foo46",
				"/foo47",
				"/foo48",
				"/foo49")
		})

		//      TODO: Convert these tests when s3 inputs are supported.
		//// in11 is an S3 input
		//in11 := client.NewS3PFSInput("", dataRepo, "")
		//in11.Pfs.Commit = commit.ID
		//t.Run("PlainS3", func(t *testing.T) {
		//	s3itr, err := NewIterator(c, in11)
		//	require.NoError(t, err)
		//	validateDI(t, s3itr, "/")

		//	// Check that every datum has an S3 input
		//	s3itr, _ = NewIterator(c, in11)
		//	var checked, s3Count int
		//	for s3itr.Next() {
		//		checked++
		//		require.Equal(t, 1, len(s3itr.Datum()))
		//		if s3itr.Datum()[0].S3 {
		//			s3Count++
		//			break
		//		}
		//	}
		//	require.True(t, checked > 0 && checked == s3Count,
		//		"checked: %v, s3Count: %v", checked, s3Count)
		//})

		//// in12 is a cross that contains an S3 input and two non-s3 inputs
		//in12 := client.NewCrossInput(in1, in2, in11)
		//t.Run("S3MixedCross", func(t *testing.T) {
		//	s3CrossItr, err := NewIterator(c, in12)
		//	require.NoError(t, err)
		//	validateDI(t, s3CrossItr,
		//		"/foo11/foo12/", "/foo21/foo12/", "/foo31/foo12/", "/foo41/foo12/",
		//		"/foo11/foo2/", "/foo21/foo2/", "/foo31/foo2/", "/foo41/foo2/",
		//		"/foo11/foo22/", "/foo21/foo22/", "/foo31/foo22/", "/foo41/foo22/",
		//		"/foo11/foo32/", "/foo21/foo32/", "/foo31/foo32/", "/foo41/foo32/",
		//		"/foo11/foo42/", "/foo21/foo42/", "/foo31/foo42/", "/foo41/foo42/",
		//	)

		//	s3CrossItr, _ = NewIterator(c, in12)
		//	var checked, s3Count int
		//	for s3CrossItr.Next() {
		//		checked++
		//		for _, d := range s3CrossItr.Datum() {
		//			if d.S3 {
		//				s3Count++
		//			}
		//		}
		//	}
		//	require.True(t, checked > 0 && checked == s3Count,
		//		"checked: %v, s3Count: %v", checked, s3Count)
		//})

		//// in13 is a cross consisting of exclusively S3 inputs
		//in13 := client.NewCrossInput(in11, in11, in11)
		//t.Run("S3OnlyCrossUnionJoin", func(t *testing.T) {
		//	s3CrossItr, err := NewIterator(c, in13)
		//	require.NoError(t, err)
		//	validateDI(t, s3CrossItr, "///")

		//	s3CrossItr, _ = NewIterator(c, in13)
		//	var checked, s3Count int
		//	for s3CrossItr.Next() {
		//		checked++
		//		for _, d := range s3CrossItr.Datum() {
		//			if d.S3 {
		//				s3Count++
		//			}
		//		}
		//	}
		//	require.True(t, checked > 0 && 3*checked == s3Count,
		//		"checked: %v, s3Count: %v", checked, s3Count)
		//})

		return nil
	}))
}

// TestJoinOnTrailingSlash tests that the same glob pattern is used for
// extracting JoinOn and GroupBy capture groups as is used to match paths. Tests
// the fix for https://github.com/pachyderm/pachyderm/v2/issues/5365
// TODO: The trailing slash glob replace is not capturing the right key and the PFS path stuff
// might need some work because it does not seem to make sense that we return a file for a path that
// ends in a trailing slash.
// Make work with V2.
//func TestJoinTrailingSlash(t *testing.T) {
//	t.Parallel()
//	db := dbutil.NewTestDB(t)
//	require.NoError(t, testpachd.WithRealEnv(db, func(env *testpachd.RealEnv) error {
//		c := env.PachClient
//		repo := []string{ // singular name b/c we only refer to individual elements
//			tu.UniqueString(t.Name() + "_0"),
//			tu.UniqueString(t.Name() + "_1"),
//		}
//		input := []*pps.Input{ // singular name b/c only use individual elements
//			client.NewPFSInputOpts("", repo[0],
//				/* commit--set below */ "", "/*", "$1", "", false, false, nil),
//			client.NewPFSInputOpts("", repo[1],
//				/* commit--set below */ "", "/*", "$1", "", false, false, nil),
//		}
//		require.NoError(t, c.CreateRepo(repo[0]))
//		require.NoError(t, c.CreateRepo(repo[1]))
//
//		// put files in structured in a way so that there are many ways to glob it
//		for i := 0; i < 2; i++ {
//			commit, err := c.StartCommit(repo[i], "master")
//			require.NoError(t, err)
//			for j := 0; j < 10; j++ {
//				require.NoError(t, c.PutFile(repo[i], commit.ID, fmt.Sprintf("foo-%v", j), strings.NewReader("bar")))
//			}
//			require.NoError(t, c.FinishCommit(repo[i], commit.ID))
//			input[i].Pfs.Commit = commit.ID
//		}
//
//		// Test without trailing slashes
//		input[0].Pfs.Glob = "/(*)"
//		input[1].Pfs.Glob = "/(*)"
//		itr, err := NewIterator(c, client.NewJoinInput(input...))
//		require.NoError(t, err)
//		validateDI(t, itr,
//			"/foo-0/foo-0",
//			"/foo-1/foo-1",
//			"/foo-2/foo-2",
//			"/foo-3/foo-3",
//			"/foo-4/foo-4",
//			"/foo-5/foo-5",
//			"/foo-6/foo-6",
//			"/foo-7/foo-7",
//			"/foo-8/foo-8",
//			"/foo-9/foo-9",
//		)
//		// Test with trailing slashes
//		input[0].Pfs.Glob = "/(*)/"
//		input[1].Pfs.Glob = "/(*)/"
//		itr, err = NewIterator(c, client.NewJoinInput(input...))
//		require.NoError(t, err)
//		validateDI(t, itr,
//			"/foo-0/foo-0",
//			"/foo-1/foo-1",
//			"/foo-2/foo-2",
//			"/foo-3/foo-3",
//			"/foo-4/foo-4",
//			"/foo-5/foo-5",
//			"/foo-6/foo-6",
//			"/foo-7/foo-7",
//			"/foo-8/foo-8",
//			"/foo-9/foo-9",
//		)
//		return nil
//	}))
//}

func validateDI(t testing.TB, di Iterator, datums ...string) {
	t.Helper()
	require.NoError(t, di.Iterate(func(meta *Meta) error {
		require.Equal(t, datums[0], computeKey(meta))
		datums = datums[1:]
		return nil
	}))
	require.Equal(t, 0, len(datums))
}

func computeKey(meta *Meta) string {
	var key string
	for _, input := range meta.Inputs {
		key += input.FileInfo.File.Path
	}
	return key
}