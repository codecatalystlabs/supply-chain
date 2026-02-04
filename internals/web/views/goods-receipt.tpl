{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Goods Receipt</h6>
  </div>
  <div class="card-body">
    <table class="windows-table" id="goodsReceiptTable">
      <thead>
        <tr>
          <th>ID</th>
          <th>System Code</th>
          <th>Facility Code</th>
          <th>Receipt Number</th>
          <th>Warehouse Ref</th>
          <th>Order Number</th>
          <th>Product Code</th>
          <th>Batch Number</th>
          <th>Quantity</th>
          <th>Receipt Date</th>
          <th>Supplier Code</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody id="goodsReceiptBody">
        <tr>
          <td colspan="12" class="empty-state">
            <i class="icon ion-ios-refresh" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
            Loading...
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</div>
{{end}}

{{define "extra_js"}}
<script>
$(document).ready(function() {
  loadGoodsReceipts();
});

function loadGoodsReceipts() {
  $.ajax({
    url: '/api/v1/goods-receipt',
    method: 'GET',
    success: function(data) {
      const tbody = $('#goodsReceiptBody');
      tbody.empty();
      
      if (data && data.length > 0) {
        data.forEach(function(item) {
          const row = `
            <tr>
              <td>${item.id}</td>
              <td>${item.grn_system_code || '-'}</td>
              <td>${item.grn_facility_code || '-'}</td>
              <td>${item.grn_facility_receipt_number || '-'}</td>
              <td>${item.grn_warehouse_ref_number || '-'}</td>
              <td>${item.grn_order_number || '-'}</td>
              <td>${item.grn_product_code || '-'}</td>
              <td>${item.grn_batch_number || '-'}</td>
              <td>${item.grn_quantity || 0}</td>
              <td>${item.grn_receipt_date ? new Date(item.grn_receipt_date).toLocaleDateString() : '-'}</td>
              <td>${item.grn_supplier_code || '-'}</td>
              <td><span class="badge badge-${item.validation_status === 1 ? 'success' : 'warning'}">${item.validation_status === 1 ? 'Valid' : 'Pending'}</span></td>
            </tr>
          `;
          tbody.append(row);
        });
      } else {
        tbody.append(`
          <tr>
            <td colspan="12" class="empty-state">
              <i class="icon ion-ios-information-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
              No goods receipts found
            </td>
          </tr>
        `);
      }
    },
    error: function(xhr, status, error) {
      $('#goodsReceiptBody').html(`
        <tr>
          <td colspan="12" class="empty-state">
            <i class="icon ion-ios-close-circle-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
            Error loading data: ${error}
          </td>
        </tr>
      `);
      toastr.error('Failed to load goods receipts data');
    }
  });
}
</script>
{{end}}
