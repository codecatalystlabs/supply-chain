{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header d-flex justify-content-between align-items-center">
    <h6 class="card-title mb-0">Role Management</h6>
  </div>
  <div class="card-body">
    <table class="windows-table" id="rolesTable">
      <thead>
        <tr>
          <th>ID</th>
          <th>Name</th>
          <th>Display Name</th>
          <th>Description</th>
          <th>Status</th>
          <th>Permissions</th>
        </tr>
      </thead>
      <tbody id="rolesBody">
        <tr>
          <td colspan="6" class="empty-state">
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
  loadRoles();
});

function loadRoles() {
  $.ajax({
    url: '/api/v1/roles',
    method: 'GET',
    success: function(data) {
      const tbody = $('#rolesBody');
      tbody.empty();
      
      if (data && data.length > 0) {
        data.forEach(function(role) {
          const permissionsList = role.permissions && role.permissions.length > 0 
            ? role.permissions.map(p => p.display_name || p.name).join(', ')
            : 'No permissions';
          const row = `
            <tr>
              <td>${role.id}</td>
              <td><strong>${role.name || '-'}</strong></td>
              <td>${role.display_name || '-'}</td>
              <td>${role.description || '-'}</td>
              <td><span class="badge badge-${role.is_active ? 'success' : 'secondary'}">${role.is_active ? 'Active' : 'Inactive'}</span></td>
              <td><small>${permissionsList}</small></td>
            </tr>
          `;
          tbody.append(row);
        });
      } else {
        tbody.append(`
          <tr>
            <td colspan="6" class="empty-state">
              <i class="icon ion-ios-information-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
              No roles found
            </td>
          </tr>
        `);
      }
    },
    error: function(xhr, status, error) {
      $('#rolesBody').html(`
        <tr>
          <td colspan="6" class="empty-state">
            <i class="icon ion-ios-close-circle-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
            Error loading data: ${error}
          </td>
        </tr>
      `);
      toastr.error('Failed to load roles');
    }
  });
}
</script>
{{end}}

