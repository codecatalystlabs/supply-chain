{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Inventory Management</h6>
  </div>
  <div class="card-body">
    <div class="alert alert-info">
      <strong>Inventory System:</strong> Monitor and manage all stock levels across facilities
    </div>
    <table class="windows-table">
      <thead>
        <tr>
          <th>Item</th>
          <th>SKU</th>
          <th>Quantity</th>
          <th>Unit</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody>
        {{if .Data}}
          {{range .Data}}
          <tr>
            <td>{{.Name}}</td>
            <td>{{.SKU}}</td>
            <td>{{.Quantity}}</td>
            <td>{{.Unit}}</td>
            <td><span class="badge badge-success">In Stock</span></td>
          </tr>
          {{end}}
        {{else}}
          <tr><td colspan="5" class="text-center text-muted">No inventory items found</td></tr>
        {{end}}
      </tbody>
    </table>
  </div>
</div>
{{end}}
