{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Stock Dispensed</h6>
  </div>
  <div class="card-body">
    <table class="windows-table" id="stockDispensedTable">
      <thead>
        <tr>
          <th>ID</th>
          <th>System Code</th>
          <th>Facility Code</th>
          <th>Product Code</th>
          <th>Batch Number</th>
          <th>Quantity</th>
          <th>Dispense Date</th>
          <th>Patient Hash</th>
          <th>Expiry Date</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody id="stockDispensedBody">
        <tr>
          <td colspan="10" class="empty-state">
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
  loadStockDispensed();
});

function loadStockDispensed() {
  $.ajax({
    url: '/api/v1/stock/dispensed',
    method: 'GET',
    success: function(data) {
      const tbody = $('#stockDispensedBody');
      tbody.empty();
      
      if (data && data.length > 0) {
        data.forEach(function(item) {
          const row = `
            <tr>
              <td>${item.id}</td>
              <td>${item.dsp_system_code || '-'}</td>
              <td>${item.dsp_facility_code || '-'}</td>
              <td>${item.dsp_product_code || '-'}</td>
              <td>${item.dsp_batch_number || '-'}</td>
              <td>${item.dsp_dispensed_quantity || 0}</td>
              <td>${item.dsp_dispense_date ? new Date(item.dsp_dispense_date).toLocaleDateString() : '-'}</td>
              <td>${item.dsp_patient_hash ? item.dsp_patient_hash.substring(0, 8) + '...' : '-'}</td>
              <td>${item.dsp_expiry_date ? new Date(item.dsp_expiry_date).toLocaleDateString() : '-'}</td>
              <td><span class="badge badge-${item.validation_status === 1 ? 'success' : 'warning'}">${item.validation_status === 1 ? 'Valid' : 'Pending'}</span></td>
            </tr>
          `;
          tbody.append(row);
        });
      } else {
        tbody.append(`
          <tr>
            <td colspan="10" class="empty-state">
              <i class="icon ion-ios-information-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
              No stock dispensed records found
            </td>
          </tr>
        `);
      }
    },
    error: function(xhr, status, error) {
      $('#stockDispensedBody').html(`
        <tr>
          <td colspan="10" class="empty-state">
            <i class="icon ion-ios-close-circle-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
            Error loading data: ${error}
          </td>
        </tr>
      `);
      toastr.error('Failed to load stock dispensed data');
    }
  });
}
</script>
{{end}}

