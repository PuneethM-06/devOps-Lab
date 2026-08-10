# Interview.md — Day 70

# Helm Templating

## 1. What is Helm templating, and what is the purpose of `{{ }}`?

**Expected Answer:**

Helm templating is the process of using template expressions and logic to dynamically generate final Kubernetes manifests.

`{{ }}` is Helm's template expression syntax. It tells Helm to evaluate whatever is inside the brackets during template rendering.

Example:

```yaml
replicas: {{ .Values.replicaCount }}
```

If `replicaCount` is `3`, Helm renders:

```yaml
replicas: 3
```

Kubernetes receives the rendered YAML, not the Helm template syntax.

---

## 2. What is the purpose of `if` / `else` in Helm templates?

**Expected Answer:**

`if` / `else` provides conditional logic in Helm templates.

It allows Helm to render different parts of a Kubernetes manifest depending on whether a condition is true or false.

Example:

```yaml
{{- if eq .Values.app_env "dev" }}
docker_image: myapp-dev
{{- else }}
docker_image: myapp-prod
{{- end }}
```

If `app_env` is `dev`, the first block is rendered. Otherwise, the `else` block is rendered.

---

## 3. What is the difference between `with` and `range`?

**Expected Answer:**

### `with`

`with` changes the current context to a nested object, making templates cleaner.

Without `with`:

```yaml
{{ .Values.image.repository }}
{{ .Values.image.tag }}
```

With `with`:

```yaml
{{- with .Values.image }}
repository: {{ .repository }}
tag: {{ .tag }}
{{- end }}
```

Inside the `with` block, `.` refers to `.Values.image`.

### `range`

`range` is used to iterate through a collection such as a list or map.

Example:

```yaml
ports:
  - 8080
  - 9090
```

Template:

```yaml
{{- range .Values.ports }}
- port: {{ . }}
{{- end }}
```

Helm processes each item in the collection.

---

## 4. What are `define`, `include`, and `_helpers.tpl` used for?

**Expected Answer:**

* `define` → Creates a reusable named template.
* `include` → Renders a named template where it is called.
* `_helpers.tpl` → Commonly stores reusable named templates/helpers.

Example:

```yaml
{{- define "myapp.labels" }}
app: myapp
team: platform
{{- end }}
```

Then:

```yaml
{{ include "myapp.labels" . }}
```

can reuse that template in another Kubernetes manifest.

The main benefit is **avoiding duplication** across Deployment, Service, Ingress, and other templates.

---

## 5. What is the purpose of `default` and `quote`?

**Expected Answer:**

### `default`

`default` provides a fallback value when the supplied value is empty or not provided.

Example:

```yaml
replicas: {{ default 3 .Values.replicaCount }}
```

If `replicaCount` is provided as `5`:

```yaml
replicas: 5
```

If it is missing or empty:

```yaml
replicas: 3
```

### `quote`

`quote` wraps a value in quotes.

Example:

```yaml
{{ .Values.appName | quote }}
```

If:

```yaml
appName: myapp
```

Helm produces:

```yaml
"myapp"
```

The `|` is the Helm pipeline operator, which passes the value to the next function.

---

# Additional Helm Templating Concepts

## `nindent`

`nindent` is used to add a newline and indent generated content.

Example:

```yaml
{{ include "myapp.labels" . | nindent 4 }}
```

This is commonly used when inserting reusable template output into YAML while maintaining the correct indentation.

---

## `indent`

`indent` adds a specified number of spaces to the generated content.

For example:

```yaml
{{ include "myapp.labels" . | indent 4 }}
```

The main purpose of `indent` and `nindent` is to keep the generated YAML correctly formatted and structured.

---

# Important Helm Template Syntax

```text
{{ .Values.x }}
```

Access a value.

```text
{{ if CONDITION }}
...
{{ end }}
```

Conditional rendering.

```text
{{ if CONDITION }}
...
{{ else }}
...
{{ end }}
```

Conditional rendering with an alternative.

```text
{{ with .Values.x }}
...
{{ end }}
```

Change the current context.

```text
{{ range .Values.x }}
...
{{ end }}
```

Loop through a collection.

```text
{{ default VALUE .Values.x }}
```

Provide a fallback value.

```text
{{ .Values.x | quote }}
```

Pass a value through the `quote` function.

```text
{{ define "name" }}
...
{{ end }}
```

Define a reusable named template.

```text
{{ include "name" . }}
```

Render a named template.

---

# Core Mental Model

```text
                    Helm Chart
                        │
                 ┌──────┴──────┐
                 ↓             ↓
             Values         Templates
                 │             │
                 │      ┌──────┴───────────┐
                 │      │                  │
                 │     if                 range
                 │     with               default
                 │     include            quote
                 │
                 └──────────┬─────────────┘
                            ↓
                     Helm Rendering
                            ↓
                  Kubernetes Manifests
                            ↓
                     Kubernetes API
                            ↓
                   Kubernetes Resources
```

# Key Takeaways

* Helm templates allow Kubernetes manifests to become reusable and dynamic.
* `{{ }}` is used for Helm template expressions.
* `.Values` allows templates to consume configuration.
* `if` / `else` provides conditional rendering.
* `with` changes the current context to a nested object.
* `range` loops through collections.
* `default` provides fallback values.
* `quote` quotes values.
* `define` creates reusable named templates.
* `include` renders reusable named templates.
* `_helpers.tpl` commonly stores reusable template definitions.
* `indent` and `nindent` help maintain correct YAML indentation.
* Helm renders the templates before Kubernetes receives the manifests.
