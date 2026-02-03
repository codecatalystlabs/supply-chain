{{template "base.tpl" .}}

{{define "breadcrumb"}}
<ol class="breadcrumb slim-breadcrumb">
  <li class="breadcrumb-item"><a href="/">Home</a></li>
  <li class="breadcrumb-item active" aria-current="page">Role Management</li>
</ol>
{{end}}

{{define "main_content"}}
<style>
.compact-table .table {
  margin-bottom: 0;
}
.compact-table .table th,
.compact-table .table td {
  padding: 0.5rem;
  font-size: 0.875rem;
}
.compact-table .btn {
  padding: 0.25rem 0.5rem;
  font-size: 0.75rem;
}
.status-badge {
  font-size: 0.75rem;
  padding: 0.25rem 0.5rem;
}
.page-header {
  margin-bottom: 1rem;
}
.page-header h4 {
  margin: 0;
}
</style>

<div class="compact-table">
  <div class="card">
    <div class="card-header page-header d-flex justify-content-between align-items-center">
      <h4>Role Management</h4>
      <a href="/cp/admin/roles/new" class="btn btn-primary">
        <i class="fa fa-plus"></i> Add New Role
      </a>
    </div>
    <div class="card-body p-0">
      {{if .error}}
      <div class="alert alert-danger m-3" role="alert">
        {{.error}}
      </div>
      {{end}}
      
      {{if .roles}}
      <div class="table-responsive">
        <table class="table table-hover mb-0">
          <thead class="thead-light">
            <tr>
              <th>Name</th>
              <th>Description</th>
              <th>Status</th>
              <th>Type</th>
              <th>Created</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {{range .roles}}
            <tr>
              <td>
                <strong>{{.Name}}</strong>
              </td>
              <td>{{.Description}}</td>
              <td>
                {{if .IsActive}}
                  <span class="badge badge-success status-badge">Active</span>
                {{else}}
                  <span class="badge badge-secondary status-badge">Inactive</span>
                {{end}}
              </td>
              <td>
                {{if .IsSystem}}
                  <span class="badge badge-warning status-badge">System</span>
                {{else}}
                  <span class="badge badge-info status-badge">User</span>
                {{end}}
              </td>
              <td>{{.CreatedAt.Format "Jan 2, 2006"}}</td>
              <td>
                <div class="btn-group" role="group">
                  <a href="/cp/admin/roles/{{.Id}}/assign-users" class="btn btn-outline-primary" title="Assign Users">
                    <i class="fa fa-users"></i>
                  </a>
                  {{if not .IsSystem}}
                  <a href="/cp/admin/roles/{{.Id}}/edit" class="btn btn-outline-secondary" title="Edit">
                    <i class="fa fa-edit"></i>
                  </a>
                  <form method="post" action="/cp/admin/roles/{{.Id}}/delete" style="display: inline;" 
                        onsubmit="return confirm('Are you sure you want to delete this role?')">
                    <button type="submit" class="btn btn-outline-danger" title="Delete">
                      <i class="fa fa-trash"></i>
                    </button>
                  </form>
                  {{else}}
                  <span class="btn btn-outline-secondary disabled" title="System role - cannot edit">
                    <i class="fa fa-lock"></i>
                  </span>
                  {{end}}
                </div>
              </td>
            </tr>
            {{end}}
          </tbody>
        </table>
      </div>
      {{else}}
      <div class="text-center py-4">
        <i class="fa fa-key fa-3x text-muted mb-3"></i>
        <p class="text-muted">No roles found. <a href="/cp/admin/roles/new">Create the first role</a></p>
      </div>
      {{end}}
    </div>
  </div>
</div>
{{end}}
