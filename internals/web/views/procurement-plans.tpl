{{template "base.tpl" .}}

{{define "content"}}
<style>
  .plan-data-table thead th { white-space: normal; min-width: 2rem; }
  #planDisplayArea .table th { white-space: normal; }
</style>
<div class="card">
  <div class="card-header">
    <h6 class="card-title mb-0">Procurement Plans</h6>
  </div>
  <div class="card-body">
    <div class="mb-4">
      <h6 class="mb-2">View plans</h6>
      <p class="text-muted small mb-3">Choose how to filter (Region, District, Facility, or Level), select the value, then click Load.</p>

      <div class="form-row align-items-end">
        <div class="form-group col-auto">
          <label for="viewBy">View by</label>
          <select id="viewBy" class="form-control" style="min-width: 130px;">
            <option value="">— Select —</option>
            <option value="region">Region</option>
            <option value="district">District</option>
            <option value="facility">Facility</option>
            <option value="level">Level</option>
          </select>
        </div>
        <div class="form-group col-auto" id="wrapRegion" style="display:none;">
          <label for="selRegion">Region</label>
          <select id="selRegion" class="form-control" style="min-width: 160px;">
            <option value="">— Select region —</option>
          </select>
        </div>
        <div class="form-group col-auto" id="wrapDistrict" style="display:none;">
          <label for="selDistrict">District</label>
          <select id="selDistrict" class="form-control" style="min-width: 160px;">
            <option value="">— Select district —</option>
          </select>
        </div>
        <div class="form-group col-auto" id="wrapFacility" style="display:none;">
          <label for="facilitySearch">Facility</label>
          <input id="facilitySearch" type="text" class="form-control mb-1" placeholder="Type to search facility…" autocomplete="off" style="min-width: 220px;">
          <select id="selFacility" class="form-control" size="8" style="min-width: 220px; max-height: 200px;">
            <option value="">— Select facility —</option>
          </select>
        </div>
        <div class="form-group col-auto" id="wrapLevel" style="display:none;">
          <label for="selLevel">Level</label>
          <select id="selLevel" class="form-control" style="min-width: 130px;">
            <option value="">— Select level —</option>
          </select>
        </div>
        <div class="form-group col-auto">
          <label for="financialYear">Financial year</label>
          <input id="financialYear" type="text" class="form-control" placeholder="e.g. 2025/26" style="width: 100px;">
        </div>
        <div class="form-group col-auto">
          <button type="button" id="btnLoadPlans" class="btn btn-primary">Load</button>
        </div>
      </div>
    </div>

    <div id="planDisplayArea" style="display:none;"></div>
    <div id="planPlaceholder" class="text-muted text-center py-5">
      Select filters above and click Load to view plans.
    </div>
  </div>
</div>
{{end}}

{{define "extra_js"}}
<script>
(function() {
  var regions = [], districts = [], facilities = [];

  function apiGet(url, data, done) {
    $.ajax({
      url: '/api/v1' + url,
      method: 'GET',
      data: data || {},
      xhrFields: { withCredentials: true }
    })
      .done(function(d) { if (done) done(d); })
      .fail(function(xhr) {
        var msg = (xhr.responseJSON && xhr.responseJSON.error) ? xhr.responseJSON.error : 'Request failed';
        if (xhr.status === 401) msg = 'Please log in again.';
        toastr.error(msg);
      });
  }

  function loadRegions() {
    $('#selRegion').find('option:not(:first)').remove().end().append('<option value="" disabled>Loading…</option>');
    apiGet('/regions', null, function(r) {
      regions = (r && Array.isArray(r)) ? r : [];
      var $sel = $('#selRegion');
      $sel.find('option').remove();
      $sel.append($('<option value="">— Select region —</option>'));
      regions.forEach(function(v) {
        var name = v.name || v.Name || '';
        if (name) $sel.append($('<option></option>').val(name).attr('data-region-id', v.id || v.ID || '').text(name));
      });
    });
  }

  function loadDistricts(regionName, done) {
    var regionId = '';
    if (regionName && regions.length) {
      var r = regions.find(function(x) { return (x.name || x.Name) === regionName; });
      if (r) regionId = (r.id != null) ? r.id : (r.ID != null) ? r.ID : '';
    }
    if (!regionId) {
      $('#selDistrict').find('option:not(:first)').remove();
      if (done) done();
      return;
    }
    apiGet('/districts', { region_id: regionId }, function(r) {
      districts = (r && Array.isArray(r)) ? r : [];
      var $sel = $('#selDistrict');
      $sel.find('option:not(:first)').remove();
      districts.forEach(function(v) {
        var name = v.name || v.Name || '';
        if (name) $sel.append($('<option></option>').val(name).text(name));
      });
      if (done) done();
    });
  }

  function loadFacilities(region, district, done) {
    var $sel = $('#selFacility');
    $sel.find('option').remove();
    $sel.append($('<option value="" disabled>Loading…</option>'));
    apiGet('/facilities', { region: region || '', district: district || '', active: 'true' }, function(r) {
      facilities = (r && Array.isArray(r)) ? r : [];
      filterFacilityOptions();
      if (done) done();
    });
  }

  function filterFacilityOptions() {
    var q = ($('#facilitySearch').val() || '').toLowerCase().trim();
    var regionName = ($('#selRegion').val() || '').trim().toLowerCase();
    var districtName = ($('#selDistrict').val() || '').trim().toLowerCase();
    var $sel = $('#selFacility');
    var $search = $('#facilitySearch');
    if (!$sel.length || !$search.length) return;
    $sel.find('option').remove();
    $sel.append($('<option value="">— Select facility —</option>'));
    var codeKey = 'facility_code';
    var nameKey = 'facility_name';
    var list = facilities.slice();
    if (regionName) {
      list = list.filter(function(f) {
        var fr = (f.region || f.Region || '').toString().trim().toLowerCase();
        return fr === regionName || fr.indexOf(regionName) >= 0 || regionName.indexOf(fr) >= 0;
      });
    }
    if (districtName) {
      list = list.filter(function(f) {
        var fd = (f.district || f.District || '').toString().trim().toLowerCase();
        return fd === districtName || fd.indexOf(districtName) >= 0 || districtName.indexOf(fd) >= 0;
      });
    }
    if (list.length === 0 && facilities.length > 0) list = facilities.slice();
    list.forEach(function(f) {
      var code = (f[codeKey] || f.FacilityCode || '').toString();
      var name = (f[nameKey] || f.FacilityName || '').toString();
      var label = code + ' – ' + name;
      var match = !q || label.toLowerCase().indexOf(q) >= 0 ||
        code.toLowerCase().indexOf(q) >= 0 ||
        name.toLowerCase().indexOf(q) >= 0;
      if (match) {
        $sel.append($('<option></option>').val(code).attr('data-name', name).text(label));
      }
    });
  }

  function loadLevels() {
    var $sel = $('#selLevel');
    if ($sel.find('option').length > 1) return;
    $sel.append($('<option value="" disabled>Loading…</option>'));
    apiGet('/levels-of-care', null, function(r) {
      var list = (r && Array.isArray(r)) ? r : [];
      $sel.find('option').remove();
      $sel.append($('<option value="">— Select level —</option>'));
      list.forEach(function(v) {
        var name = v.name || v.Name || '';
        if (name) $sel.append($('<option></option>').val(name).text(name));
      });
    });
  }

  $('#viewBy').on('change', function() {
    var v = $(this).val();
    $('#wrapRegion, #wrapDistrict, #wrapFacility, #wrapLevel').hide();
    $('#selRegion, #selDistrict, #selFacility').val('');
    $('#facilitySearch').val('');
    if (v === 'region') {
      $('#wrapRegion').show();
      loadRegions();
    } else if (v === 'district') {
      $('#wrapRegion').show();
      $('#wrapDistrict').show();
      loadRegions();
      loadDistricts($('#selRegion').val());
    } else if (v === 'facility') {
      $('#wrapRegion').show();
      $('#wrapDistrict').show();
      $('#wrapFacility').show();
      $('#selFacility').find('option').remove();
      $('#selFacility').append($('<option value="">— Select facility —</option>'));
      $('#facilitySearch').val('');
      loadRegions();
      loadFacilities('', '', null);
    } else if (v === 'level') {
      $('#wrapLevel').show();
      loadLevels();
    }
  });

  $('#selRegion').on('change', function() {
    var r = $(this).val();
    loadDistricts(r);
    if ($('#viewBy').val() === 'facility') {
      $('#selDistrict').val('');
      filterFacilityOptions();
    }
  });

  $('#selDistrict').on('change', function() {
    if ($('#viewBy').val() !== 'facility') return;
    filterFacilityOptions();
  });

  $(document).on('input keyup', '#facilitySearch', function() {
    if ($('#viewBy').val() !== 'facility') return;
    filterFacilityOptions();
  });

  function buildQueryParams() {
    var viewBy = $('#viewBy').val(), fy = $('#financialYear').val().trim();
    var params = {};
    if (fy) params.financial_year = fy;
    if (!viewBy) return params;
    if (viewBy === 'region') params.region = $('#selRegion').val() || '';
    else if (viewBy === 'district') params.district = $('#selDistrict').val() || '';
    else if (viewBy === 'facility') params.facility_id = $('#selFacility').val() || '';
    else if (viewBy === 'level') params.level_of_care = $('#selLevel').val() || '';
    return params;
  }

  $('#btnLoadPlans').on('click', function() {
    var viewBy = $('#viewBy').val();
    if (!viewBy) {
      toastr.warning('Select "View by" first.');
      return;
    }
    var params = buildQueryParams();
    var need = (viewBy === 'region' && !params.region) || (viewBy === 'district' && !params.district) || (viewBy === 'facility' && !params.facility_id) || (viewBy === 'level' && !params.level_of_care);
    if (need) {
      toastr.warning('Please select a ' + viewBy + ' value.');
      return;
    }
    $('#planPlaceholder').hide();
    $('#planDisplayArea').show().html('<p class="text-center py-4">Loading…</p>');
    $.ajax({
      url: '/api/v1/procurement-plans',
      method: 'GET',
      data: params,
      xhrFields: { withCredentials: true },
      success: function(data) {
        renderPlanDisplay(data || []);
      },
      error: function(xhr) {
        $('#planDisplayArea').hide();
        $('#planPlaceholder').show();
        toastr.error(xhr.responseJSON && xhr.responseJSON.error ? xhr.responseJSON.error : 'Failed to load plans');
      }
    });
  });

  function fmtNum(v) {
    if (v == null || v === '') return '-';
    var n = Number(v);
    return isNaN(n) ? v : n.toLocaleString();
  }
  function fmtPct(v) {
    if (v == null || v === '') return '-';
    var n = Number(v);
    return isNaN(n) ? v : n + '%';
  }

  function renderPlanDisplay(data) {
    if (!data.length) {
      $('#planDisplayArea').html('<p class="text-center text-muted py-4">No procurement plans found.</p>');
      return;
    }
    var byFac = {};
    data.forEach(function(row) {
      var key = (row.facility_id || '') + '|' + (row.financial_year || '');
      if (!byFac[key]) byFac[key] = { header: row, rows: [] };
      byFac[key].rows.push(row);
    });

    var html = '';
    $.each(byFac, function(key, group) {
      var h = group.header;
      var fy = h.financial_year || '';
      var title = 'FY' + fy + ' GENERAL HOSPITAL PROCUREMENT PLANNING FORM';
      var venRows = {};
      var totalSpend = 0;
      group.rows.forEach(function(r) {
        var ven = (r.ven_classification || 'N').toUpperCase();
        if (ven !== 'V' && ven !== 'E') ven = 'N';
        var price = r.unit_price != null ? r.unit_price : 0;
        var qty = r.bi_monthly_plan_qty != null ? r.bi_monthly_plan_qty : (r.previous_bi_monthly_planned_qty != null ? r.previous_bi_monthly_planned_qty : 0);
        var lineTotal = price * (typeof qty === 'number' ? qty : 0);
        if (!venRows[ven]) venRows[ven] = 0;
        venRows[ven] += lineTotal;
        totalSpend += lineTotal;
      });
      var venV = venRows['V'] || 0, venE = venRows['E'] || 0, venN = venRows['N'] || 0;
      var pctV = totalSpend ? ((venV / totalSpend) * 100).toFixed(0) : 0;
      var pctE = totalSpend ? ((venE / totalSpend) * 100).toFixed(0) : 0;
      var pctN = totalSpend ? ((venN / totalSpend) * 100).toFixed(0) : 0;
      var indAnn = (h.indicative_annual_budget != null) ? h.indicative_annual_budget : 0;
      var calcAnn = (h.calculated_annual_procurement != null) ? h.calculated_annual_procurement : totalSpend;
      var indBi = (h.indicative_bi_monthly_budget != null) ? h.indicative_bi_monthly_budget : 0;
      var calcBi = (h.calculated_bi_monthly_procurement != null) ? h.calculated_bi_monthly_procurement : (totalSpend / 6);
      var remain = (h.remaining_budget != null) ? h.remaining_budget : (indAnn - calcAnn);
      var pctRemain = (h.percent_budget_remaining != null) ? h.percent_budget_remaining : (indAnn ? ((remain / indAnn) * 100).toFixed(1) : 0);

      html += '<div class="card mb-4 border">';
      html += '<div class="card-body">';
      html += '<h6 class="text-danger font-weight-bold text-center mb-3">' + title + '</h6>';
      html += '<div class="row mb-3">';
      html += '<div class="col-md-4"><table class="table table-sm table-bordered"><tbody>';
      html += '<tr><td><strong>Name</strong></td><td>' + (h.facility_name || '-') + '</td></tr>';
      html += '<tr><td><strong>Code</strong></td><td>' + (h.facility_code || h.facility_id || '-') + '</td></tr>';
      html += '<tr><td><strong>Level</strong></td><td>' + (h.level_of_care || '-') + '</td></tr>';
      html += '<tr><td><strong>District</strong></td><td>' + (h.district || '-') + '</td></tr>';
      html += '<tr><td><strong>Region</strong></td><td>' + (h.region || '-') + '</td></tr>';
      html += '<tr><td><strong>Zone</strong></td><td>' + (h.zone || '-') + '</td></tr>';
      html += '</tbody></table></div>';
      html += '<div class="col-md-4"><table class="table table-sm table-bordered"><thead><tr><th>VEN CLASSIFICATION SUMMARY</th><th>V</th><th>E</th><th>N</th></tr></thead><tbody>';
      html += '<tr><td>Spend in UGX</td><td>' + fmtNum(venV) + '</td><td>' + fmtNum(venE) + '</td><td>' + fmtNum(venN) + '</td></tr>';
      html += '<tr><td>% of Budget</td><td>' + pctV + '%</td><td>' + pctE + '%</td><td>' + pctN + '%</td></tr>';
      html += '</tbody></table></div>';
      html += '<div class="col-md-4"><table class="table table-sm table-bordered"><tbody>';
      html += '<tr><td>Indicative Annual Planning Budget</td><td>' + fmtNum(indAnn) + '</td></tr>';
      html += '<tr><td>Calculated Annual Procurement</td><td>' + fmtNum(calcAnn) + '</td></tr>';
      html += '<tr><td>Indicative Bi-monthly Planning Budget</td><td>' + fmtNum(indBi) + '</td></tr>';
      html += '<tr><td>Calculated Bi-Monthly Procurement</td><td>' + fmtNum(calcBi) + '</td></tr>';
      html += '<tr><td>Remaining Budget</td><td>' + fmtNum(remain) + '</td></tr>';
      html += '<tr><td>% Budget Remaining</td><td>' + fmtPct(pctRemain) + '</td></tr>';
      html += '</tbody></table></div>';
      html += '</div>';

      html += '<div class="table-responsive"><table class="windows-table plan-data-table">';
      html += '<thead><tr>';
      var cols = ['CODE','DESCRIPTION','UNIT','VEN','PRICE','FY2024/25 BIMONTHLY PLANNED QTY','PAST AVG. NMS ISSUE PLAN Q1','2024 ADJUSTED AMC','FY2025/26 BIMONTHLY PLAN QTY','TOTAL','COMMENT','FUNDED QTY'];
      cols.forEach(function(c) { html += '<th>' + c + '</th>'; });
      html += '</tr><tr class="filter-row">';
      for (var i = 0; i < cols.length; i++) { html += '<th><input type="text" class="form-control form-control-sm col-filter" data-col="' + i + '" placeholder="Filter"></th>'; }
      html += '</tr></thead><tbody>';
      group.rows.forEach(function(r) {
        var price = r.unit_price != null ? r.unit_price : 0;
        var qty = r.bi_monthly_plan_qty != null ? r.bi_monthly_plan_qty : (r.previous_bi_monthly_planned_qty != null ? r.previous_bi_monthly_planned_qty : 0);
        var total = price * (typeof qty === 'number' ? qty : 0);
        html += '<tr>';
        html += '<td>' + (r.product_code || '-') + '</td>';
        html += '<td>' + (r.product_description || '-') + '</td>';
        html += '<td>' + (r.unit_of_measure || '-') + '</td>';
        html += '<td>' + (r.ven_classification || '-') + '</td>';
        html += '<td>' + fmtNum(r.unit_price) + '</td>';
        html += '<td>' + (r.previous_bi_monthly_planned_qty != null ? r.previous_bi_monthly_planned_qty : '-') + '</td>';
        html += '<td>' + (r.past_avg_nms_issue_plan_qty != null ? fmtNum(r.past_avg_nms_issue_plan_qty) : '-') + '</td>';
        html += '<td>' + (r.adjusted_amc != null ? fmtNum(r.adjusted_amc) : '-') + '</td>';
        html += '<td>' + (r.bi_monthly_plan_qty != null ? r.bi_monthly_plan_qty : '-') + '</td>';
        html += '<td>' + fmtNum(total) + '</td>';
        html += '<td>' + (r.comment || '-') + '</td>';
        html += '<td>' + (r.funded_qty != null ? r.funded_qty : '-') + '</td>';
        html += '</tr>';
      });
      html += '</tbody></table></div>';
      html += '</div></div>';
    });
    $('#planDisplayArea').html(html);

    $('.plan-data-table').each(function() {
      var $t = $(this);
      $t.find('.col-filter').on('input', function() {
        var col = $(this).data('col');
        var val = $(this).val().toLowerCase();
        $t.find('tbody tr').each(function() {
          var cell = $(this).find('td').eq(col).text().toLowerCase();
          $(this).toggle(!val || cell.indexOf(val) >= 0);
        });
      });
    });
  }

  $(document).ready(function() {
    loadRegions();
  });
})();
</script>
{{end}}
