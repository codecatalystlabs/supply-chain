{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Stock on Hand</h6>
  </div>
  <div class="card-body">
    <table class="table table-hover table-striped">
      <thead>
        <tr>
          <th>Product</th>
          <th>Warehouse</th>
          <th>Quantity</th>
          <th>Unit Cost</th>
          <th>Total Value</th>
          <th>Last Updated</th>
        </tr>
      </thead>
      <tbody>
        {{if .Data}}
          {{range .Data}}
          <tr>
            <td>{{.ProductName}}</td>
            <td>{{.WarehouseName}}</td>
            <td>{{.Quantity}}</td>
            <td>{{.UnitCost}}</td>
            <td>{{.TotalValue}}</td>
            <td>{{.UpdatedAt}}</td>
          </tr>
          {{end}}
        {{else}}
          <tr><td colspan="6" class="text-center text-muted">No stock found</td></tr>
        {{end}}
      </tbody>
    </table>
  </div>
</div>
{{end}}
