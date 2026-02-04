{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Stock Adjustments</h6>
  </div>
  <div class="card-body">
    <table class="windows-table" id="stockAdjustmentsTable">
      <thead>
        <tr>
          <th>ID</th>
          <th>System Code</th>
          <th>Facility Code</th>
          <th>Pharmacy</th>
          <th>Product Code</th>
          <th>Batch Number</th>
          <th>Adjustment Type</th>
          <th>Quantity</th>
          <th>Adjustment Date</th>
          <th>Reason</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody id="stockAdjustmentsBody">
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
  loadStockAdjustments();
});

function loadStockAdjustments() {
  $.ajax({
    url: '/api/v1/stock/adjustment',
    method: 'GET',
    success: function(data) {
      const tbody = $('#stockAdjustmentsBody');
      tbody.empty();
      
      if (data && data.length > 0) {
        data.forEach(function(item) {
          const adjustmentTypeBadge = item.adj_adjustment_type === 'increase' ? 'success' : 'danger';
          const row = `
            <tr>
              <td>${item.id}</td>
              <td>${item.adj_system_code || '-'}</td>
              <td>${item.adj_facility_code || '-'}</td>
              <td>${item.adj_pharmacy_name || item.adj_pharmacy_code || '-'}</td>
              <td>${item.adj_product_code || '-'}</td>
              <td>${item.adj_batch_number || '-'}</td>
              <td><span class="badge badge-${adjustmentTypeBadge}">${item.adj_adjustment_type || '-'}</span></td>
              <td>${item.adj_quantity || 0}</td>
              <td>${item.adj_adjustment_date ? new Date(item.adj_adjustment_date).toLocaleDateString() : '-'}</td>
              <td>${item.adj_adjustment_reason || '-'}</td>
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
              No stock adjustment records found
            </td>
          </tr>
        `);
      }
    },
    error: function(xhr, status, error) {
      $('#stockAdjustmentsBody').html(`
        <tr>
          <td colspan="11" class="empty-state">
            <i class="icon ion-ios-close-circle-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
            Error loading data: ${error}
          </td>
        </tr>
      `);
      toastr.error('Failed to load stock adjustments data');
    }
  });
}
</script>
{{end}}

