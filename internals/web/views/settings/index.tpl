{{template "base.tpl" .}}

{{define "breadcrumb"}}
<ol class="breadcrumb slim-breadcrumb">
  <li class="breadcrumb-item"><a href="/">Home</a></li>
  <li class="breadcrumb-item active" aria-current="page">Settings</li>
</ol>
{{end}}

{{define "main_content"}}
<style>
.settings-container {
  max-width: 1000px;
  margin: 0 auto;
}
.settings-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
  margin-bottom: 2rem;
}
@media (max-width: 768px) {
  .settings-row {
    grid-template-columns: 1fr;
    gap: 1rem;
  }
}
.settings-card {
  border: 1px solid #dee2e6;
  border-radius: 0.25rem;
}
.settings-card-header {
  background-color: #f8f9fa;
  border-bottom: 1px solid #dee2e6;
  padding: 1rem;
}
.settings-card-header h5 {
  margin: 0;
  font-weight: 600;
}
.settings-card-body {
  padding: 1.5rem;
}
.form-group label {
  font-weight: 500;
  margin-bottom: 0.5rem;
}
.form-control-static {
  padding-top: 0.5rem;
  padding-bottom: 0.5rem;
}
.password-requirements {
  background-color: #e8f4f8;
  border-left: 4px solid #0c5460;
  padding: 0.75rem;
  margin-top: 1rem;
  font-size: 0.875rem;
  color: #0c5460;
}
.password-requirements ul {
  margin-bottom: 0;
  padding-left: 1.5rem;
}
.password-requirements li {
  margin-bottom: 0.25rem;
}
.password-input-group {
  position: relative;
  display: flex;
  align-items: center;
}
.password-input-group .form-control {
  padding-right: 2.5rem;
}
.password-toggle-btn {
  position: absolute;
  right: 0.75rem;
  background: none;
  border: none;
  cursor: pointer;
  color: #6c757d;
  padding: 0.25rem 0.5rem;
  font-size: 1rem;
  transition: color 0.2s ease;
}
.password-toggle-btn:hover {
  color: #495057;
}
.password-toggle-btn:focus {
  outline: none;
}
</style>

<div class="settings-container">
  {{if .error}}
  <div class="alert alert-danger alert-dismissible fade show" role="alert">
    {{.error}}
    <button type="button" class="close" data-dismiss="alert" aria-label="Close">
      <span aria-hidden="true">&times;</span>
    </button>
  </div>
  {{end}}

  {{if .successMessage}}
  <div class="alert alert-success alert-dismissible fade show" role="alert">
    {{.successMessage}}
    <button type="button" class="close" data-dismiss="alert" aria-label="Close">
      <span aria-hidden="true">&times;</span>
    </button>
  </div>
  {{end}}

  <div class="settings-row">
    <!-- User Information Card -->
    <div class="settings-card">
      <div class="settings-card-header">
        <h5><i class="fa fa-user"></i> User Information</h5>
      </div>
      <div class="settings-card-body">
        <div class="form-group">
          <label>Username</label>
          <p class="form-control-static">{{.user.Username}}</p>
        </div>
        {{if .user.Contact}}
        <div class="form-group">
          <label>Contact</label>
          <p class="form-control-static">{{.user.Contact}}</p>
        </div>
        {{end}}
        <div class="form-group">
          <label>Account Status</label>
          <p class="form-control-static">
            {{if .user.IsActive}}
            <span class="badge badge-success">Active</span>
            {{else}}
            <span class="badge badge-danger">Inactive</span>
            {{end}}
          </p>
        </div>
      </div>
    </div>

    <!-- Change Password Card -->
    <div class="settings-card">
      <div class="settings-card-header">
        <h5><i class="fa fa-lock"></i> Change Password</h5>
      </div>
      <div class="settings-card-body">
        <form method="post" action="/cp/settings/change-password">
          <div class="form-group">
            <label for="current_password">Current Password *</label>
            <div class="password-input-group">
              <input type="password" class="form-control password-input" id="current_password" name="current_password" required>
              <button type="button" class="password-toggle-btn" onclick="togglePasswordVisibility(this)">
                <i class="fa fa-eye"></i>
              </button>
            </div>
            <small class="form-text text-muted">Enter your current password</small>
          </div>

          <div class="form-group">
            <label for="new_password">New Password *</label>
            <div class="password-input-group">
              <input type="password" class="form-control password-input" id="new_password" name="new_password" required>
              <button type="button" class="password-toggle-btn" onclick="togglePasswordVisibility(this)">
                <i class="fa fa-eye"></i>
              </button>
            </div>
            <small class="form-text text-muted">Enter a new password</small>
          </div>

          <div class="form-group">
            <label for="confirm_password">Confirm Password *</label>
            <div class="password-input-group">
              <input type="password" class="form-control password-input" id="confirm_password" name="confirm_password" required>
              <button type="button" class="password-toggle-btn" onclick="togglePasswordVisibility(this)">
                <i class="fa fa-eye"></i>
              </button>
            </div>
            <small class="form-text text-muted">Confirm your new password</small>
          </div>

          <div class="password-requirements">
            <strong>Password Requirements:</strong>
            <ul>
              <li>At least 6 characters long</li>
              <li>Must be different from current password</li>
              <li>Both new password fields must match</li>
            </ul>
          </div>

          <div class="form-group mt-3">
            <button type="submit" class="btn btn-primary">
              <i class="fa fa-save"></i> Change Password
            </button>
            <a href="/cp/home" class="btn btn-secondary ml-2">
              <i class="fa fa-times"></i> Cancel
            </a>
          </div>
        </form>
      </div>
    </div>
  </div>
</div>

<script>
function togglePasswordVisibility(button) {
  const inputGroup = button.closest('.password-input-group');
  const input = inputGroup.querySelector('.password-input');
  const icon = button.querySelector('i');
  
  if (input.type === 'password') {
    input.type = 'text';
    icon.classList.remove('fa-eye');
    icon.classList.add('fa-eye-slash');
  } else {
    input.type = 'password';
    icon.classList.remove('fa-eye-slash');
    icon.classList.add('fa-eye');
  }
}
</script>
{{end}}
