{{ define "Index" }}
  {{ template "Header" }}
    {{ template "Menu"  }}
    <h2> Registered </h2>
    <table border="1">
      <thead>
      <tr>
        <td>ID</td>
        <td>Name</td>
        <td>City</td>
        <td>View</td>
        <td>Edit</td>
        <td>Delete</td>
      </tr>
       </thead>
       <tbody>
    {{ range . }}
      <tr>
        <td>{{ .Id }}</td>
        <td> {{ .Name }} </td>
        <td>{{ .City }} </td> 
        <td><a href="/show?id={{ .Id }}">View</a></td>
        <td><a href="/edit?id={{ .Id }}">Edit</a></td>
        <td><a href="/delete?id={{ .Id }}">Delete</a><td>
      </tr>
    {{ end }}
       </tbody>
    </table>
  {{ template "Footer" }}
{{ end }}