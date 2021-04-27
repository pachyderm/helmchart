# Helm Chart Style Guide

Due to the nature of helm, there are many ways to accomplish the same goal. There are also sometimes unintended side effects when chosing a particular style over another.  

## Be explicit 

When checking whether to include or exclude a given manifest, use explicit values instead of inferring what the template should do. This makes it clearer to the user what will happen and reduce unintended side effects.

Instead of:

```
{{- if .Values.dash.url -}}
apiVersion: "extensions/v1beta1"
kind: "Ingress"
metadata:
  name: "dash"
...
```

Use:

```
{{- if .Values.dash.enabled -}}
apiVersion: "extensions/v1beta1"
kind: "Ingress"
metadata:
  name: "dash"
...
```

## Allow for future extensibility via Nested Values

By using nested values, you can add new values without breaking backward compatibility. The caveat is that you must ensure that you are always setting at least one value in the nested object, or you must check for the existence of the object in the chart.

Nested vs Flat Values

Nested:

```
server:
  name: nginx
  port: 80
```
Flat:

```
serverName: nginx
serverPort: 80
```

## Build in extensibility

For annotations, node affinities and other optional configuration use `toYaml` to allow for greater customization.

```
annotations: {{ toYaml .Values.ingress.annotations | nindent 4 }}
```

## Prefer `nindent` over `indent`
It can be difficult to line up indentation using `indent` as it includes spaces which preceed the template.

```
annotations: 
  {{ toYaml .Values.ingress.annotations | indent 4 }}
```


```
annotations: {{ toYaml .Values.ingress.annotations | nindent 4 }}
```

## Further Reading:

- [The Art of Helm Chart Patterns](https://hackernoon.com/the-art-of-the-helm-chart-patterns-from-the-official-kubernetes-charts-8a7cafa86d12)
- [Official Best Practices](https://helm.sh/docs/chart_best_practices/)
- [Helm from Basics to Advanced](https://banzaicloud.com/blog/creating-helm-charts/)
