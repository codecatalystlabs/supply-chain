{{template "base.tpl" .}}

{{define "breadcrumb"}}
<ol class="breadcrumb slim-breadcrumb">
  <li class="breadcrumb-item"><a href="/">Home</a></li>
  <li class="breadcrumb-item"><a href="/cp/admin/roles">Role Management</a></li>
  <li class="breadcrumb-item active" aria-current="page">Add New Role</li>
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
.compact-form select.form-control {
  padding: 0.25rem 0.5rem;
  font-size: 0.875rem;
  height: auto;
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
.permissions-section {
  max-height: 300px;
  overflow-y: auto;
  border: 1px solid #dee2e6;
  border-radius: 0.25rem;
  padding: 0.75rem;
}
.permission-group {
  margin-bottom: 1rem;
}
.permission-group h6 {
  color: #495057;
  font-weight: 600;
  margin-bottom: 0.5rem;
  padding-bottom: 0.25rem;
  border-bottom: 1px solid #e9ecef;
}
.permission-item {
  margin-bottom: 0.25rem;
}
.permission-item .form-check {
  margin-bottom: 0;
}
.permission-item .form-check-label {
  font-size: 0.875rem;
}
</style>

<div class="row">
  <div class="col-md-10 offset-md-1">
    <div class="card compact-form">
      <div class="card-header">
        <h5 class="card-title mb-0">Add New Role</h5>
      </div>
      <div class="card-body">
        {{if .error}}
        <div class="alert alert-danger" role="alert">
          {{.error}}
        </div>
        {{end}}
        
        <form method="post" action="/cp/admin/roles/create">
          <div class="form-group">
            <label for="name">Role Name *</label>
            <input type="text" class="form-control" id="name" name="name" 
                   value="{{if .role}}{{.role.Name}}{{end}}" required>
            <small class="form-text text-muted">Unique name for the role (e.g., "Manager", "Viewer")</small>
          </div>
          
          <div class="form-group">
            <label for="description">Description</label>
            <textarea class="form-control" id="description" name="description" rows="2">{{if .role}}{{.role.Description}}{{end}}</textarea>
            <small class="form-text text-muted">Brief description of what this role can do</small>
          </div>
          
          <div class="form-group">
            <div class="form-check">
              <input type="checkbox" class="form-check-input" id="is_active" name="is_active" 
                     {{if .role}}{{if .role.IsActive}}checked{{end}}{{else}}checked{{end}}>
              <label class="form-check-label" for="is_active">
                Active
              </label>
            </div>
            <small class="form-text text-muted">Whether this role is currently active and can be assigned</small>
          </div>
          
          <div class="form-group">
            <label>Permissions</label>
            <div class="permissions-section">
              {{if .permissions}}
                {{$currentResource := ""}}
                {{range .permissions}}
                  {{if ne .Resource $currentResource}}
                    {{if ne $currentResource ""}}
                      </div> <!-- End previous permission group -->
                    {{end}}
                    {{$currentResource = .Resource}}
                    <div class="permission-group">
                      <h6>{{.Resource}} Permissions</h6>
                  {{end}}
                  <div class="permission-item">
                    <div class="form-check">
                      <input type="checkbox" class="form-check-input" id="perm_{{.Id}}" 
                             name="permission_ids" value="{{.Id}}">
                      <label class="form-check-label" for="perm_{{.Id}}">
                        <strong>{{.Action}}</strong> - {{.Description}}
                      </label>
                    </div>
                  </div>
                {{end}}
                {{if ne $currentResource ""}}
                  </div> <!-- End last permission group -->
                {{end}}
              {{else}}
                <p class="text-muted">No permissions available</p>
              {{end}}
            </div>
            <small class="form-text text-muted">Select the permissions this role should have</small>
          </div>
          
          <div class="form-group mt-2">
            <button type="submit" class="btn btn-primary">
              <i class="fa fa-save"></i> Create Role
            </button>
            <a href="/cp/admin/roles" class="btn btn-secondary ml-2">
              <i class="fa fa-times"></i> Cancel
            </a>
          </div>
        </form>
      </div>
    </div>
  </div>
</div>
{{end}}
