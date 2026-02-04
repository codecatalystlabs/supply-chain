{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Pharmacy Stock</h6>
  </div>
  <div class="card-body">
    <table class="windows-table">
      <thead>
        <tr>
          <th>Item</th>
          <th>Pharmacy</th>
          <th>Quantity</th>
          <th>Expiry Date</th>
          <th>Status</th>
          <th>Last Updated</th>
        </tr>
      </thead>
      <tbody>
        {{if .Data}}
          {{range .Data}}
          <tr>
            <td>{{.ItemName}}</td>
            <td>{{.PharmacyName}}</td>
            <td>{{.Quantity}}</td>
            <td>{{.ExpiryDate}}</td>
            <td><span class="badge badge-success">In Stock</span></td>
            <td>{{.UpdatedAt}}</td>
          </tr>
          {{end}}
        {{else}}
          <tr><td colspan="6" class="text-center text-muted">No pharmacy stock found</td></tr>
        {{end}}
      </tbody>
    </table>
  </div>
</div>
{{end}}
