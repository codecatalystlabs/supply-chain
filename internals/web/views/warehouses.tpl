{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Warehouses</h6>
  </div>
  <div class="card-body">
    <table class="windows-table" id="warehousesTable">
      <thead>
        <tr>
          <th>ID</th>
          <th>Code</th>
          <th>Name</th>
          <th>Type</th>
          <th>Address</th>
          <th>Contact Person</th>
          <th>Status</th>
          <th>Created</th>
        </tr>
      </thead>
      <tbody id="warehousesBody">
        <tr>
          <td colspan="8" class="empty-state">
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
  loadWarehouses();
});

function loadWarehouses() {
  $.ajax({
    url: '/api/v1/warehouses',
    method: 'GET',
    success: function(data) {
      const tbody = $('#warehousesBody');
      tbody.empty();
      
      if (data && data.length > 0) {
        data.forEach(function(item) {
          const row = `
            <tr>
              <td>${item.id}</td>
              <td>${item.warehouse_code || '-'}</td>
              <td>${item.warehouse_name || '-'}</td>
              <td>${item.warehouse_type || '-'}</td>
              <td>${item.address || '-'}</td>
              <td>${item.contact_person || '-'}</td>
              <td><span class="badge badge-${item.is_active ? 'success' : 'secondary'}">${item.is_active ? 'Active' : 'Inactive'}</span></td>
              <td>${item.created_at ? new Date(item.created_at).toLocaleDateString() : '-'}</td>
            </tr>
          `;
          tbody.append(row);
        });
      } else {
        tbody.append(`
          <tr>
            <td colspan="8" class="empty-state">
              <i class="icon ion-ios-information-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
              No warehouses found
            </td>
          </tr>
        `);
      }
    },
    error: function(xhr, status, error) {
      $('#warehousesBody').html(`
        <tr>
          <td colspan="8" class="empty-state">
            <i class="icon ion-ios-close-circle-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
            Error loading data: ${error}
          </td>
        </tr>
      `);
      toastr.error('Failed to load warehouses data');
    }
  });
}
</script>
{{end}}
