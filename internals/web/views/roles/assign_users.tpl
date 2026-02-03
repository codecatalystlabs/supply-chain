{{template "base.tpl" .}}

{{define "breadcrumb"}}
<ol class="breadcrumb slim-breadcrumb">
  <li class="breadcrumb-item"><a href="/">Home</a></li>
  <li class="breadcrumb-item"><a href="/cp/admin/roles">Role Management</a></li>
  <li class="breadcrumb-item active" aria-current="page">Assign Users to {{.role.Name}}</li>
</ol>
{{end}}

{{define "main_content"}}
<style>
.compact-form .form-group {
  margin-bottom: 0.75rem;
}
.compact-form .form-control {
  padding: 0.25rem 0.5rem;
  font-size: 0.875rem;
}
.compact-form .form-text {
  font-size: 0.75rem;
  margin-top: 0.125rem;
}
.compact-form label {
  margin-bottom: 0.25rem;
  font-weight: 500;
}
.compact-form .card-body {
  padding: 1rem;
}
.compact-form .btn {
  padding: 0.375rem 0.75rem;
  font-size: 0.875rem;
}
.compact-form .card-header {
  padding: 0.75rem 1rem;
}
.compact-form .alert {
  padding: 0.5rem 0.75rem;
  margin-bottom: 0.75rem;
}
.compact-form .form-check-label {
  margin-bottom: 0;
}
.users-section {
  max-height: 400px;
  overflow-y: auto;
  border: 1px solid #dee2e6;
  border-radius: 0.25rem;
  padding: 0.75rem;
}
.user-item {
  margin-bottom: 0.5rem;
  padding: 0.5rem;
  border: 1px solid #e9ecef;
  border-radius: 0.25rem;
  background-color: #f8f9fa;
}
.user-item:last-child {
  margin-bottom: 0;
}
.user-item .form-check {
  margin-bottom: 0;
}
.user-item .form-check-label {
  font-size: 0.875rem;
  width: 100%;
}
.user-info {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.user-details {
  flex-grow: 1;
}
.user-status {
  font-size: 0.75rem;
}
.select-all-section {
  margin-bottom: 1rem;
  padding: 0.5rem;
  background-color: #e7f3ff;
  border-radius: 0.25rem;
}
</style>

<div class="row">
  <div class="col-md-10 offset-md-1">
    <div class="card compact-form">
      <div class="card-header">
        <h5 class="card-title mb-0">Assign Users to Role: {{.role.Name}}</h5>
      </div>
      <div class="card-body">
        {{if .error}}
        <div class="alert alert-danger" role="alert">
          {{.error}}
        </div>
        {{end}}
        
        <div class="alert alert-info">
          <strong>Role:</strong> {{.role.Name}}<br>
          <strong>Description:</strong> {{.role.Description}}
        </div>
        
        <form method="post" action="/cp/admin/roles/{{.role.Id}}/update-assignments">
          <div class="form-group">
            <label>Select Users for this Role</label>
            
            {{if .users}}
            <div class="select-all-section">
              <div class="form-check">
                <input type="checkbox" class="form-check-input" id="select_all" onchange="toggleAllUsers()">
                <label class="form-check-label" for="select_all">
                  <strong>Select/Deselect All Users</strong>
                </label>
              </div>
            </div>
            
            <div class="users-section">
              {{$assignedUsers := .assigned_users}}
              {{range .users}}
              <div class="user-item">
                <div class="form-check">
                  <input type="checkbox" class="form-check-input user-checkbox" id="user_{{.Id}}" 
                         name="user_ids" value="{{.Id}}"
                         {{if index $assignedUsers .Id}}checked{{end}}>
                  <label class="form-check-label" for="user_{{.Id}}">
                    <div class="user-info">
                      <div class="user-details">
                        <strong>{{.Username}}</strong>
                        {{if .Contact}}<br><small class="text-muted">{{.Contact}}</small>{{end}}
                      </div>
                      <div class="user-status">
                        {{if .IsActive}}
                          <span class="badge badge-success">Active</span>
                        {{else}}
                          <span class="badge badge-secondary">Inactive</span>
                        {{end}}
                        {{if index $assignedUsers .Id}}
                          <span class="badge badge-primary">Currently Assigned</span>
                        {{end}}
                      </div>
                    </div>
                  </label>
                </div>
              </div>
              {{end}}
            </div>
            {{else}}
            <div class="text-center py-4">
              <i class="fa fa-users fa-3x text-muted mb-3"></i>
              <p class="text-muted">No users found in the system</p>
            </div>
            {{end}}
            
            <small class="form-text text-muted">
              Select the users who should have this role. Users can have multiple roles.
            </small>
          </div>
          
          <div class="form-group mt-2">
            <button type="submit" class="btn btn-primary">
              <i class="fa fa-save"></i> Update Role Assignments
            </button>
            <a href="/cp/admin/roles" class="btn btn-secondary ml-2">
              <i class="fa fa-arrow-left"></i> Back to Roles
            </a>
          </div>
        </form>
      </div>
    </div>
  </div>
</div>

<script>
function toggleAllUsers() {
  const selectAllCheckbox = document.getElementById('select_all');
  const userCheckboxes = document.querySelectorAll('.user-checkbox');
  
  userCheckboxes.forEach(checkbox => {
    checkbox.checked = selectAllCheckbox.checked;
  });
}

// Update select all checkbox based on individual selections
document.addEventListener('DOMContentLoaded', function() {
  const selectAllCheckbox = document.getElementById('select_all');
  const userCheckboxes = document.querySelectorAll('.user-checkbox');
  
  function updateSelectAll() {
    const totalCheckboxes = userCheckboxes.length;
    const checkedCheckboxes = document.querySelectorAll('.user-checkbox:checked').length;
    
    if (checkedCheckboxes === 0) {
      selectAllCheckbox.checked = false;
      selectAllCheckbox.indeterminate = false;
    } else if (checkedCheckboxes === totalCheckboxes) {
      selectAllCheckbox.checked = true;
      selectAllCheckbox.indeterminate = false;
    } else {
      selectAllCheckbox.checked = false;
      selectAllCheckbox.indeterminate = true;
    }
  }
  
  // Add event listeners to individual checkboxes
  userCheckboxes.forEach(checkbox => {
    checkbox.addEventListener('change', updateSelectAll);
  });
  
  // Initialize select all state
  updateSelectAll();
});
</script>
{{end}}
