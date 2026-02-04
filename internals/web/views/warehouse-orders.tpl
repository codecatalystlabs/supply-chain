{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Warehouse Orders</h6>
  </div>
  <div class="card-body">
    <table class="windows-table">
      <thead>
        <tr>
          <th>Order ID</th>
          <th>Source Warehouse</th>
          <th>Destination</th>
          <th>Items Count</th>
          <th>Date</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody>
        {{if .Data}}
          {{range .Data}}
          <tr>
            <td>{{.OrderID}}</td>
            <td>{{.SourceWarehouse}}</td>
            <td>{{.DestinationWarehouse}}</td>
            <td>{{.ItemsCount}}</td>
            <td>{{.OrderDate}}</td>
            <td><span class="badge badge-info">{{.Status}}</span></td>
          </tr>
          {{end}}
        {{else}}
          <tr><td colspan="6" class="text-center text-muted">No warehouse orders found</td></tr>
        {{end}}
      </tbody>
    </table>
  </div>
</div>
{{end}}
