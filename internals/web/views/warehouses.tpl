{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Warehouses</h6>
  </div>
  <div class="card-body">
    <table class="windows-table">
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
            <td><span class="status-badge active">Active</span></td>
          </tr>
          {{end}}
        {{else}}
          <tr><td colspan="6" class="empty-state"><i class="icon ion-ios-information-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>No warehouses found</td></tr>
        {{end}}
      </tbody>
    </table>
  </div>
</div>
{{end}}
