{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Warehouses</h6>
  </div>
  <div class="card-body">
    <table class="table table-hover table-striped">
      <thead>
        <tr>
          <th>ID</th>
          <th>Name</th>
          <th>Location</th>
          <th>Capacity</th>
          <th>Current Stock</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody>
        {{if .Data}}
          {{range .Data}}
          <tr>
            <td>{{.ID}}</td>
            <td>{{.Name}}</td>
            <td>{{.Location}}</td>
            <td>{{.Capacity}}</td>
            <td>{{.CurrentStock}}</td>
            <td><span class="badge badge-success">Active</span></td>
          </tr>
          {{end}}
        {{else}}
          <tr><td colspan="6" class="text-center text-muted">No warehouses found</td></tr>
        {{end}}
      </tbody>
    </table>
  </div>
</div>
{{end}}
