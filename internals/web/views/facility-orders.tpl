{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Facility Orders</h6>
  </div>
  <div class="card-body">
    <table class="windows-table" id="facilityOrdersTable">
      <thead>
        <tr>
          <th>ID</th>
          <th>Order Number</th>
          <th>Order Ref</th>
          <th>Facility Code</th>
          <th>Warehouse Code</th>
          <th>Order Date</th>
          <th>Order Type</th>
          <th>Status</th>
          <th>Total Items</th>
          <th>Total Quantity</th>
          <th>Priority</th>
        </tr>
      </thead>
      <tbody id="facilityOrdersBody">
        <tr>
          <td colspan="11" class="empty-state">
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
  loadFacilityOrders();
});

function loadFacilityOrders() {
  $.ajax({
    url: '/api/v1/facility-orders',
    method: 'GET',
    success: function(data) {
      const tbody = $('#facilityOrdersBody');
      tbody.empty();
      
      if (data && data.length > 0) {
        data.forEach(function(item) {
          const statusBadge = getOrderStatusBadge(item.order_status);
          const row = `
            <tr>
              <td>${item.id}</td>
              <td>${item.order_number || '-'}</td>
              <td>${item.order_ref_number || '-'}</td>
              <td>${item.facility_code || '-'}</td>
              <td>${item.warehouse_code || '-'}</td>
              <td>${item.order_date ? new Date(item.order_date).toLocaleDateString() : '-'}</td>
              <td>${item.order_type || '-'}</td>
              <td>${statusBadge}</td>
              <td>${item.total_items || 0}</td>
              <td>${item.total_quantity || 0}</td>
              <td>${item.priority || '-'}</td>
            </tr>
          `;
          tbody.append(row);
        });
      } else {
        tbody.append(`
          <tr>
            <td colspan="11" class="empty-state">
              <i class="icon ion-ios-information-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
              No facility orders found
            </td>
          </tr>
        `);
      }
    },
    error: function(xhr, status, error) {
      $('#facilityOrdersBody').html(`
        <tr>
          <td colspan="11" class="empty-state">
            <i class="icon ion-ios-close-circle-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
            Error loading data: ${error}
          </td>
        </tr>
      `);
      toastr.error('Failed to load facility orders data');
    }
  });
}

function getOrderStatusBadge(status) {
  const badges = {
    'draft': '<span class="badge badge-secondary">Draft</span>',
    'submitted': '<span class="badge badge-info">Submitted</span>',
    'approved': '<span class="badge badge-success">Approved</span>',
    'rejected': '<span class="badge badge-danger">Rejected</span>',
    'delivered': '<span class="badge badge-primary">Delivered</span>'
  };
  return badges[status] || '<span class="badge badge-secondary">' + (status || 'Unknown') + '</span>';
}
</script>
{{end}}
