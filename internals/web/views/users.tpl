{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header d-flex justify-content-between align-items-center">
    <h6 class="card-title mb-0">User Management</h6>
    <button class="btn btn-primary btn-sm" onclick="showCreateUserModal()">
      <i class="fa fa-plus"></i> Add User
    </button>
  </div>
  <div class="card-body">
    <table class="windows-table" id="usersTable">
      <thead>
        <tr>
          <th>ID</th>
          <th>Username</th>
          <th>Email</th>
          <th>Full Name</th>
          <th>Facility</th>
          <th>Roles</th>
          <th>Status</th>
          <th>Last Login</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody id="usersBody">
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

<!-- Create/Edit User Modal -->
<div class="modal fade" id="userModal" tabindex="-1" role="dialog">
  <div class="modal-dialog modal-lg" role="document">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title" id="userModalTitle">Create User</h5>
        <button type="button" class="close" data-dismiss="modal">
          <span>&times;</span>
        </button>
      </div>
      <div class="modal-body">
        <form id="userForm">
          <input type="hidden" id="userId" name="id">
          <div class="form-group">
            <label>Username *</label>
            <input type="text" class="form-control" id="username" name="username" required>
          </div>
          <div class="form-group">
            <label>Email *</label>
            <input type="email" class="form-control" id="email" name="email" required>
          </div>
          <div class="form-group">
            <label>First Name *</label>
            <input type="text" class="form-control" id="firstName" name="first_name" required>
          </div>
          <div class="form-group">
            <label>Last Name *</label>
            <input type="text" class="form-control" id="lastName" name="last_name" required>
          </div>
          <div class="form-group" id="passwordGroup">
            <label>Password *</label>
            <input type="password" class="form-control" id="password" name="password" minlength="6">
            <small class="form-text text-muted">Leave blank to keep existing password</small>
          </div>
          <div class="form-group">
            <label>Facility</label>
            <select class="form-control" id="facilityId" name="facility_id">
              <option value="">Select Facility</option>
            </select>
          </div>
          <div class="form-group">
            <label>Roles</label>
            <div id="rolesCheckboxes"></div>
          </div>
          <div class="form-group">
            <div class="form-check">
              <input type="checkbox" class="form-check-input" id="isActive" name="is_active" checked>
              <label class="form-check-label" for="isActive">Active</label>
            </div>
          </div>
        </form>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-secondary" data-dismiss="modal">Cancel</button>
        <button type="button" class="btn btn-primary" onclick="saveUser()">Save</button>
      </div>
    </div>
  </div>
</div>
{{end}}

{{define "extra_js"}}
<script>
let facilities = [];
let roles = [];
let editingUserId = null;

$(document).ready(function() {
  loadUsers();
  loadFacilities();
  loadRoles();
});

function loadUsers() {
  $.ajax({
    url: '/api/v1/users',
    method: 'GET',
    success: function(data) {
      const tbody = $('#usersBody');
      tbody.empty();
      
      if (data && data.length > 0) {
        data.forEach(function(user) {
          const rolesList = user.roles && user.roles.length > 0 
            ? user.roles.map(r => r.display_name).join(', ')
            : 'No roles';
          const row = `
            <tr>
              <td>${user.id}</td>
              <td>${user.username || '-'}</td>
              <td>${user.email || '-'}</td>
              <td>${user.first_name} ${user.last_name}</td>
              <td>${user.facility_name || '-'}</td>
              <td>${rolesList}</td>
              <td><span class="badge badge-${user.is_active ? 'success' : 'secondary'}">${user.is_active ? 'Active' : 'Inactive'}</span></td>
              <td>${user.last_login_at ? new Date(user.last_login_at).toLocaleString() : 'Never'}</td>
              <td>
                <button class="btn btn-sm btn-info" onclick="editUser(${user.id})">
                  <i class="fa fa-edit"></i>
                </button>
                <button class="btn btn-sm btn-danger" onclick="deleteUser(${user.id})">
                  <i class="fa fa-trash"></i>
                </button>
              </td>
            </tr>
          `;
          tbody.append(row);
        });
      } else {
        tbody.append(`
          <tr>
            <td colspan="9" class="empty-state">
              <i class="icon ion-ios-information-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
              No users found
            </td>
          </tr>
        `);
      }
    },
    error: function(xhr, status, error) {
      $('#usersBody').html(`
        <tr>
          <td colspan="9" class="empty-state">
            <i class="icon ion-ios-close-circle-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
            Error loading data: ${error}
          </td>
        </tr>
      `);
      toastr.error('Failed to load users');
    }
  });
}

function loadFacilities() {
  $.ajax({
    url: '/api/v1/facilities',
    method: 'GET',
    success: function(data) {
      facilities = data || [];
      const select = $('#facilityId');
      select.empty().append('<option value="">Select Facility</option>');
      facilities.forEach(function(facility) {
        select.append(`<option value="${facility.id}">${facility.facility_name}</option>`);
      });
    }
  });
}

function loadRoles() {
  $.ajax({
    url: '/api/v1/roles',
    method: 'GET',
    success: function(data) {
      roles = data || [];
      const container = $('#rolesCheckboxes');
      container.empty();
      roles.forEach(function(role) {
        container.append(`
          <div class="form-check">
            <input type="checkbox" class="form-check-input role-checkbox" value="${role.id}" id="role_${role.id}">
            <label class="form-check-label" for="role_${role.id}">${role.display_name}</label>
          </div>
        `);
      });
    },
    error: function() {
      // If roles endpoint doesn't exist, create a simple list
      const container = $('#rolesCheckboxes');
      container.html('<p class="text-muted">Roles will be loaded from API</p>');
    }
  });
}

function showCreateUserModal() {
  editingUserId = null;
  $('#userModalTitle').text('Create User');
  $('#userForm')[0].reset();
  $('#userId').val('');
  $('#passwordGroup').show();
  $('#password').attr('required', true);
  $('#userModal').modal('show');
}

function editUser(id) {
  editingUserId = id;
  $.ajax({
    url: '/api/v1/users/' + id,
    method: 'GET',
    success: function(user) {
      $('#userModalTitle').text('Edit User');
      $('#userId').val(user.id);
      $('#username').val(user.username).prop('disabled', true);
      $('#email').val(user.email);
      $('#firstName').val(user.first_name);
      $('#lastName').val(user.last_name);
      $('#facilityId').val(user.facility_id || '');
      $('#isActive').prop('checked', user.is_active);
      $('#passwordGroup').hide();
      $('#password').removeAttr('required');
      
      // Set roles
      $('.role-checkbox').prop('checked', false);
      if (user.roles) {
        user.roles.forEach(function(role) {
          $('#role_' + role.id).prop('checked', true);
        });
      }
      
      $('#userModal').modal('show');
    },
    error: function() {
      toastr.error('Failed to load user');
    }
  });
}

function saveUser() {
  const formData = {
    username: $('#username').val(),
    email: $('#email').val(),
    first_name: $('#firstName').val(),
    last_name: $('#lastName').val(),
    facility_id: $('#facilityId').val() ? parseInt($('#facilityId').val()) : null,
    is_active: $('#isActive').is(':checked'),
    role_ids: $('.role-checkbox:checked').map(function() {
      return parseInt($(this).val());
    }).get()
  };

  if (!editingUserId) {
    // Create
    formData.password = $('#password').val();
    $.ajax({
      url: '/api/v1/users',
      method: 'POST',
      contentType: 'application/json',
      data: JSON.stringify(formData),
      success: function() {
        toastr.success('User created successfully');
        $('#userModal').modal('hide');
        loadUsers();
      },
      error: function(xhr) {
        const error = xhr.responseJSON ? xhr.responseJSON.error : 'Failed to create user';
        toastr.error(error);
      }
    });
  } else {
    // Update
    if ($('#password').val()) {
      toastr.warning('Password changes not supported via this interface');
    }
    $.ajax({
      url: '/api/v1/users/' + editingUserId,
      method: 'PUT',
      contentType: 'application/json',
      data: JSON.stringify(formData),
      success: function() {
        toastr.success('User updated successfully');
        $('#userModal').modal('hide');
        loadUsers();
      },
      error: function(xhr) {
        const error = xhr.responseJSON ? xhr.responseJSON.error : 'Failed to update user';
        toastr.error(error);
      }
    });
  }
}

function deleteUser(id) {
  if (!confirm('Are you sure you want to delete this user?')) {
    return;
  }
  $.ajax({
    url: '/api/v1/users/' + id,
    method: 'DELETE',
    success: function() {
      toastr.success('User deleted successfully');
      loadUsers();
    },
    error: function() {
      toastr.error('Failed to delete user');
    }
  });
}
</script>
{{end}}

