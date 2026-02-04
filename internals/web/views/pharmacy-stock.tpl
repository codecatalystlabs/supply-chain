{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Pharmacy Stock</h6>
  </div>
  <div class="card-body">
    <table class="windows-table" id="pharmacyStockTable">
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
      <tbody id="pharmacyStockBody">
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
  loadPharmacyStock();
});

function loadPharmacyStock() {
  $.ajax({
    url: '/api/v1/pharmacy-stock',
    method: 'GET',
    success: function(data) {
      const tbody = $('#pharmacyStockBody');
      tbody.empty();
      
      if (data && data.length > 0) {
        data.forEach(function(item) {
          const row = `
            <tr>
              <td>${item.id}</td>
              <td>${item.pha_system_code || '-'}</td>
              <td>${item.pha_facility_code || '-'}</td>
              <td>${item.pha_product_code || '-'}</td>
              <td>${item.pha_batch_number || '-'}</td>
              <td>${item.pha_quantity || 0}</td>
              <td>${item.pha_expiry_date ? new Date(item.pha_expiry_date).toLocaleDateString() : '-'}</td>
              <td>${item.pha_timestamp ? new Date(item.pha_timestamp).toLocaleString() : '-'}</td>
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
              No pharmacy stock found
            </td>
          </tr>
        `);
      }
    },
    error: function(xhr, status, error) {
      $('#pharmacyStockBody').html(`
        <tr>
          <td colspan="9" class="empty-state">
            <i class="icon ion-ios-close-circle-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
            Error loading data: ${error}
          </td>
        </tr>
      `);
      toastr.error('Failed to load pharmacy stock data');
    }
  });
}
</script>
{{end}}
