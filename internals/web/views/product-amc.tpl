{{template "base.tpl" .}}

{{define "content"}}
<style>
  /* Windows-like professional table styling */
  .windows-table {
    border-collapse: collapse;
    width: 100%;
    font-size: 13px;
    background: #fff;
    box-shadow: 0 2px 4px rgba(0,0,0,0.15), 0 1px 2px rgba(0,0,0,0.1);
  }
  
  .windows-table thead {
    background: linear-gradient(to bottom, #f8f9fa 0%, #e9ecef 100%);
    border-bottom: 2px solid #dee2e6;
  }
  
  .windows-table thead th {
    padding: 6px 12px;
    font-weight: 600;
    font-size: 12px;
    color: #495057;
    text-transform: uppercase;
    letter-spacing: 0.3px;
    border-right: 1px solid #dee2e6;
    white-space: nowrap;
    text-align: left;
  }
  
  .windows-table thead th:last-child {
    border-right: none;
  }
  
  .windows-table tbody tr {
    border-bottom: 1px solid #e9ecef;
    transition: background-color 0.15s ease;
  }
  
  .windows-table tbody tr:hover {
    background-color: #e7f3ff;
  }
  
  .windows-table tbody tr:nth-child(even) {
    background-color: #fafbfc;
  }
  
  .windows-table tbody tr:nth-child(even):hover {
    background-color: #e7f3ff;
  }
  
  .windows-table tbody td {
    padding: 7px 12px;
    color: #212529;
    border-right: 1px solid #f0f0f0;
    vertical-align: middle;
  }
  
  .windows-table tbody td:last-child {
    border-right: none;
  }
  
  /* Badge styling */
  .status-badge {
    display: inline-block;
    padding: 3px 8px;
    font-size: 11px;
    font-weight: 600;
    border-radius: 3px;
    text-transform: uppercase;
    letter-spacing: 0.3px;
  }
  
  .status-badge.active {
    background-color: #d4edda;
    color: #155724;
    border: 1px solid #c3e6cb;
  }
  
  .status-badge.inactive {
    background-color: #f8d7da;
    color: #721c24;
    border: 1px solid #f5c6cb;
  }
  
  /* Card adjustments */
  .card {
    border: 1px solid #dee2e6;
    box-shadow: 0 2px 4px rgba(0,0,0,0.08);
  }
  
  .card-header {
    background: linear-gradient(to bottom, #ffffff 0%, #f8f9fa 100%);
    border-bottom: 1px solid #dee2e6;
    padding: 10px 15px;
  }
  
  .card-title {
    margin: 0;
    font-size: 14px;
    font-weight: 600;
    color: #495057;
  }
  
  .card-body {
    padding: 0;
  }
  
  .empty-state {
    text-align: center;
    padding: 40px 20px;
    color: #6c757d;
    font-size: 13px;
  }
</style>

<div class="card">
  <div class="card-header">
    <h6 class="card-title">Product AMC </h6>
  </div>
  <div class="card-body">
    <table class="windows-table">
      <thead>
        <tr>
          <th>Product</th>
          <th>Facility</th>
          <th>AMC</th>
          <th>Period</th>
          <th>Last Updated</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody>
        {{if .Data}}
          {{range .Data}}
          <tr>
            <td>{{.ProductName}}</td>
            <td>{{.FacilityName}}</td>
            <td><strong>{{.AMC}}</strong></td>
            <td>{{.Period}}</td>
            <td>{{.UpdatedAt}}</td>
            <td><span class="status-badge active">Active</span></td>
          </tr>
          {{end}}
        {{else}}
          <tr>
            <td colspan="6" class="empty-state">
              <i class="icon ion-ios-information-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
              No product AMC data found
            </td>
          </tr>
        {{end}}
      </tbody>
    </table>
  </div>
</div>
{{end}}
