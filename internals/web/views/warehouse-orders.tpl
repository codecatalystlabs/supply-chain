{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Warehouse Orders</h6>
  </div>
  <div class="card-body">
    <table class="windows-table" id="warehouseOrdersTable">
      <thead>
        <tr>
          <th>ID</th>
          <th>Warehouse Code</th>
          <th>Order Number</th>
          <th>Received Date</th>
          <th>Honored Quantity</th>
          <th>Delivered Count</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody id="warehouseOrdersBody">
        <tr>
          <td colspan="7" class="empty-state">
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
  loadWarehouseOrders();
});

function loadWarehouseOrders() {
  $.ajax({
    url: '/api/v1/warehouse-orders',
    method: 'GET',
    success: function(data) {
      const tbody = $('#warehouseOrdersBody');
      tbody.empty();
      
      if (data && data.length > 0) {
        data.forEach(function(item) {
          const statusBadge = getWarehouseOrderStatusBadge(item.status);
          const row = `
            <tr>
              <td>${item.id}</td>
              <td>${item.warehouse_code || '-'}</td>
              <td>${item.order_number || '-'}</td>
              <td>${item.received_date ? new Date(item.received_date).toLocaleDateString() : '-'}</td>
              <td>${item.honored_quantity || 0}</td>
              <td>${item.delivered_count || 0}</td>
              <td>${statusBadge}</td>
            </tr>
          `;
          tbody.append(row);
        });
      } else {
        tbody.append(`
          <tr>
            <td colspan="7" class="empty-state">
              <i class="icon ion-ios-information-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
              No warehouse orders found
            </td>
          </tr>
        `);
      }
    },
    error: function(xhr, status, error) {
      $('#warehouseOrdersBody').html(`
        <tr>
          <td colspan="7" class="empty-state">
            <i class="icon ion-ios-close-circle-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
            Error loading data: ${error}
          </td>
        </tr>
      `);
      toastr.error('Failed to load warehouse orders data');
    }
  });
}

function getWarehouseOrderStatusBadge(status) {
  const badges = {
    'pending': '<span class="badge badge-warning">Pending</span>',
    'received': '<span class="badge badge-success">Received</span>',
    'delivered': '<span class="badge badge-primary">Delivered</span>',
    'completed': '<span class="badge badge-info">Completed</span>'
  };
  return badges[status] || '<span class="badge badge-secondary">' + (status || 'Unknown') + '</span>';
}
</script>
{{end}}
