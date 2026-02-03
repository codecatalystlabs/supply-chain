{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Purchase Orders</h6>
  </div>
  <div class="card-body">
    <table class="table table-hover table-striped">
      <thead>
        <tr>
          <th>PO Number</th>
          <th>Supplier</th>
          <th>Order Date</th>
          <th>Expected Delivery</th>
          <th>Total Amount</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody>
        {{if .Data}}
          {{range .Data}}
          <tr>
            <td>{{.PONumber}}</td>
            <td>{{.SupplierName}}</td>
            <td>{{.OrderDate}}</td>
            <td>{{.ExpectedDelivery}}</td>
            <td>{{.TotalAmount}}</td>
            <td><span class="badge badge-warning">{{.Status}}</span></td>
          </tr>
          {{end}}
        {{else}}
          <tr><td colspan="6" class="text-center text-muted">No purchase orders found</td></tr>
        {{end}}
      </tbody>
    </table>
  </div>
</div>
{{end}}
