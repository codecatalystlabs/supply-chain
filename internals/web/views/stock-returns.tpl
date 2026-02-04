{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Stock Returns</h6>
  </div>
  <div class="card-body">
    <table class="windows-table" id="stockReturnsTable">
      <thead>
        <tr>
          <th>ID</th>
          <th>System Code</th>
          <th>Facility Code</th>
          <th>Return Number</th>
          <th>Return Date</th>
          <th>Product Code</th>
          <th>Batch Number</th>
          <th>Unit Code</th>
          <th>Quantity</th>
          <th>Timestamp</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody id="stockReturnsBody">
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
  loadStockReturns();
});

function loadStockReturns() {
  $.ajax({
    url: '/api/v1/stock/return',
    method: 'GET',
    success: function(data) {
      const tbody = $('#stockReturnsBody');
      tbody.empty();
      
      if (data && data.length > 0) {
        data.forEach(function(item) {
          const row = `
            <tr>
              <td>${item.id}</td>
              <td>${item.rtn_system_code || '-'}</td>
              <td>${item.rtn_facility_code || '-'}</td>
              <td>${item.rtn_return_number || '-'}</td>
              <td>${item.rtn_return_date ? new Date(item.rtn_return_date).toLocaleDateString() : '-'}</td>
              <td>${item.rtn_product_code || '-'}</td>
              <td>${item.rtn_batch_number || '-'}</td>
              <td>${item.rtn_unit_code || '-'}</td>
              <td>${item.rtn_quantity || 0}</td>
              <td>${item.rtn_timestamp ? new Date(item.rtn_timestamp).toLocaleString() : '-'}</td>
              <td><span class="badge badge-${item.validation_status === 1 ? 'success' : 'warning'}">${item.validation_status === 1 ? 'Valid' : 'Pending'}</span></td>
            </tr>
          `;
          tbody.append(row);
        });
      } else {
        tbody.append(`
          <tr>
            <td colspan="11" class="empty-state">
              <i class="icon ion-ios-information-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
              No stock return records found
            </td>
          </tr>
        `);
      }
    },
    error: function(xhr, status, error) {
      $('#stockReturnsBody').html(`
        <tr>
          <td colspan="11" class="empty-state">
            <i class="icon ion-ios-close-circle-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
            Error loading data: ${error}
          </td>
        </tr>
      `);
      toastr.error('Failed to load stock returns data');
    }
  });
}
</script>
{{end}}

