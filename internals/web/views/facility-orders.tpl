{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Facility Orders</h6>
  </div>
  <div class="card-body">
    <table class="windows-table">
      <thead>
        <tr>
          <th>Order ID</th>
          <th>Facility</th>
          <th>Order Date</th>
          <th>Status</th>
          <th>Total Items</th>
          <th>Amount</th>
        </tr>
      </thead>
      <tbody>
        {{if .Data}}
          {{range .Data}}
          <tr>
            <td>{{.OrderID}}</td>
            <td>{{.FacilityName}}</td>
            <td>{{.OrderDate}}</td>
            <td><span class="badge badge-warning">{{.Status}}</span></td>
            <td>{{.TotalItems}}</td>
            <td>{{.Amount}}</td>
          </tr>
          {{end}}
        {{else}}
          <tr>
            <td colspan="6" class="empty-state">
              <i class="icon ion-ios-information-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
              No facility orders found
            </td>
          </tr>
        {{end}}
      </tbody>
    </table>
  </div>
</div>
{{end}}
