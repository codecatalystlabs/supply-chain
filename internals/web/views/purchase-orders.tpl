{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Purchase Orders</h6>
  </div>
  <div class="card-body">
    <table class="windows-table" id="purchaseOrdersTable">
      <thead>
        <tr>
          <th>ID</th>
          <th>System Code</th>
          <th>Facility Code</th>
          <th>Order Number</th>
          <th>Order Ref</th>
          <th>Product Code</th>
          <th>Quantity</th>
          <th>Order Date</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody id="purchaseOrdersBody">
        <tr>
          <td colspan="9" class="empty-state">
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
  loadPurchaseOrders();
});

function loadPurchaseOrders() {
  $.ajax({
    url: '/api/v1/purchase-orders',
    method: 'GET',
    success: function(data) {
      const tbody = $('#purchaseOrdersBody');
      tbody.empty();
      
      if (data && data.length > 0) {
        data.forEach(function(item) {
          const row = `
            <tr>
              <td>${item.id}</td>
              <td>${item.ord_system_code || '-'}</td>
              <td>${item.ord_facility_code || '-'}</td>
              <td>${item.ord_order_number || '-'}</td>
              <td>${item.ord_order_ref_number || '-'}</td>
              <td>${item.ord_product_code || '-'}</td>
              <td>${item.ord_ordered_quantity || 0}</td>
              <td>${item.ord_order_date ? new Date(item.ord_order_date).toLocaleDateString() : '-'}</td>
              <td><span class="badge badge-${item.validation_status === 1 ? 'success' : 'warning'}">${item.validation_status === 1 ? 'Valid' : 'Pending'}</span></td>
            </tr>
          `;
          tbody.append(row);
        });
      } else {
        tbody.append(`
          <tr>
            <td colspan="9" class="empty-state">
              <i class="icon ion-ios-information-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
              No purchase orders found
            </td>
          </tr>
        `);
      }
    },
    error: function(xhr, status, error) {
      $('#purchaseOrdersBody').html(`
        <tr>
          <td colspan="9" class="empty-state">
            <i class="icon ion-ios-close-circle-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
            Error loading data: ${error}
          </td>
        </tr>
      `);
      toastr.error('Failed to load purchase orders data');
    }
  });
}
</script>
{{end}}
