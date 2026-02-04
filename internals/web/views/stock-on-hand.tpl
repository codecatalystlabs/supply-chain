{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Stock on Hand</h6>
  </div>
  <div class="card-body">
    <table class="windows-table" id="stockOnHandTable">
      <thead>
        <tr>
          <th>ID</th>
          <th>System Code</th>
          <th>Facility Code</th>
          <th>Product Code</th>
          <th>Batch Number</th>
          <th>Quantity</th>
          <th>Expiry Date</th>
          <th>Timestamp</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody id="stockOnHandBody">
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
  loadStockOnHand();
});

function loadStockOnHand() {
  $.ajax({
    url: '/api/v1/stock/on-hand',
    method: 'GET',
    success: function(data) {
      const tbody = $('#stockOnHandBody');
      tbody.empty();
      
      if (data && data.length > 0) {
        data.forEach(function(item) {
          const row = `
            <tr>
              <td>${item.id}</td>
              <td>${item.src_system_code || '-'}</td>
              <td>${item.src_facility_code || '-'}</td>
              <td>${item.src_product_code || '-'}</td>
              <td>${item.src_batch_number || '-'}</td>
              <td>${item.src_quantity || 0}</td>
              <td>${item.src_expiry_date ? new Date(item.src_expiry_date).toLocaleDateString() : '-'}</td>
              <td>${item.src_timestamp ? new Date(item.src_timestamp).toLocaleString() : '-'}</td>
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
              No stock on hand records found
            </td>
          </tr>
        `);
      }
    },
    error: function(xhr, status, error) {
      $('#stockOnHandBody').html(`
        <tr>
          <td colspan="9" class="empty-state">
            <i class="icon ion-ios-close-circle-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
            Error loading data: ${error}
          </td>
        </tr>
      `);
      toastr.error('Failed to load stock on hand data');
    }
  });
}
</script>
{{end}}

