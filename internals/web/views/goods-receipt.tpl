{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Goods Receipt</h6>
  </div>
  <div class="card-body">
    <table class="windows-table">
      <thead>
        <tr>
          <th>Receipt ID</th>
          <th>Warehouse</th>
          <th>PO Number</th>
          <th>Receipt Date</th>
          <th>Items</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody>
        {{if .Data}}
          {{range .Data}}
          <tr>
            <td>{{.ReceiptID}}</td>
            <td>{{.Warehouse}}</td>
            <td>{{.PONumber}}</td>
            <td>{{.ReceiptDate}}</td>
            <td>{{.ItemCount}}</td>
            <td><span class="badge badge-success">{{.Status}}</span></td>
          </tr>
          {{end}}
        {{else}}
          <tr>
            <td colspan="6" class="empty-state">
              <i class="icon ion-ios-information-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
              No goods receipts found
            </td>
          </tr>
        {{end}}
      </tbody>
    </table>
  </div>
</div>
{{end}}
