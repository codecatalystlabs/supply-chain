{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Stock Transfers</h6>
  </div>
  <div class="card-body">
    <table class="windows-table" id="stockTransfersTable">
      <thead>
        <tr>
          <th>ID</th>
          <th>Transfer Ref</th>
          <th>Type</th>
          <th>From Facility</th>
          <th>To Facility</th>
          <th>Product Code</th>
          <th>Batch Number</th>
          <th>Quantity</th>
          <th>Transfer Date</th>
          <th>Status</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody id="stockTransfersBody">
        <tr>
          <td colspan="11" class="empty-state">
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
  loadStockTransfers();
});

function loadStockTransfers() {
  $.ajax({
    url: '/api/v1/stock/transfers',
    method: 'GET',
    success: function(data) {
      const tbody = $('#stockTransfersBody');
      tbody.empty();
      
      if (data && data.length > 0) {
        data.forEach(function(item) {
          const statusBadge = getStatusBadge(item.status);
          const actions = getTransferActions(item);
          const row = `
            <tr>
              <td>${item.id}</td>
              <td>${item.transfer_ref || '-'}</td>
              <td>${item.transfer_type || '-'}</td>
              <td>${item.from_facility_code || '-'}</td>
              <td>${item.to_facility_code || '-'}</td>
              <td>${item.product_code || '-'}</td>
              <td>${item.batch_number || '-'}</td>
              <td>${item.quantity || 0}</td>
              <td>${item.transfer_date ? new Date(item.transfer_date).toLocaleDateString() : '-'}</td>
              <td>${statusBadge}</td>
              <td>${actions}</td>
            </tr>
          `;
          tbody.append(row);
        });
      } else {
        tbody.append(`
          <tr>
            <td colspan="11" class="empty-state">
              <i class="icon ion-ios-information-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
              No stock transfer records found
            </td>
          </tr>
        `);
      }
    },
    error: function(xhr, status, error) {
      $('#stockTransfersBody').html(`
        <tr>
          <td colspan="11" class="empty-state">
            <i class="icon ion-ios-close-circle-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
            Error loading data: ${error}
          </td>
        </tr>
      `);
      toastr.error('Failed to load stock transfers data');
    }
  });
}

function getStatusBadge(status) {
  const badges = {
    'pending': '<span class="badge badge-warning">Pending</span>',
    'approved': '<span class="badge badge-info">Approved</span>',
    'in_transit': '<span class="badge badge-primary">In Transit</span>',
    'received': '<span class="badge badge-success">Received</span>',
    'rejected': '<span class="badge badge-danger">Rejected</span>'
  };
  return badges[status] || '<span class="badge badge-secondary">' + (status || 'Unknown') + '</span>';
}

function getTransferActions(item) {
  let actions = '';
  if (item.status === 'pending') {
    actions += '<button class="btn btn-sm btn-success" onclick="approveTransfer(' + item.id + ')">Approve</button> ';
  }
  if (item.status === 'approved' || item.status === 'in_transit') {
    actions += '<button class="btn btn-sm btn-primary" onclick="receiveTransfer(' + item.id + ')">Receive</button>';
  }
  return actions || '-';
}

function approveTransfer(id) {
  $.ajax({
    url: '/api/v1/stock/transfers/' + id + '/approve',
    method: 'POST',
    success: function() {
      toastr.success('Transfer approved successfully');
      loadStockTransfers();
    },
    error: function() {
      toastr.error('Failed to approve transfer');
    }
  });
}

function receiveTransfer(id) {
  $.ajax({
    url: '/api/v1/stock/transfers/' + id + '/receive',
    method: 'POST',
    success: function() {
      toastr.success('Transfer received successfully');
      loadStockTransfers();
    },
    error: function() {
      toastr.error('Failed to receive transfer');
    }
  });
}
</script>
{{end}}

