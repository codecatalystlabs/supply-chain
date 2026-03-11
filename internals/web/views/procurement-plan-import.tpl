{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title mb-0">Import Procurement Plan</h6>
  </div>
  <div class="card-body">
    <p class="mb-4">Select Region, Zone, District, and Level of care from the lookup tables, then choose a facility and upload the XLS/XLSX file.</p>

    <form id="uploadProcurementXls" class="procurement-import-form">
      <div class="form-row">
        <div class="form-group col-md-3">
          <label for="importRegion">Region</label>
          <select id="importRegion" class="form-control">
            <option value="">Select region</option>
          </select>
        </div>
        <div class="form-group col-md-3">
          <label for="importZone">Zone</label>
          <select id="importZone" class="form-control">
            <option value="">Select zone</option>
          </select>
        </div>
        <div class="form-group col-md-3">
          <label for="importDistrict">District</label>
          <select id="importDistrict" class="form-control">
            <option value="">Select district</option>
          </select>
        </div>
        <div class="form-group col-md-3">
          <label for="importLevelOfCare">Level of care</label>
          <select id="importLevelOfCare" class="form-control">
            <option value="">Select level</option>
          </select>
        </div>
      </div>
      <div class="form-row mt-2">
        <div class="form-group col-md-4">
          <label for="importFacilitySearch">Facility (type to search)</label>
          <input id="importFacilitySearch" type="text" class="form-control" placeholder="Type to filter…" autocomplete="off">
          <select id="importFacility" class="form-control mt-2" size="5" style="max-height: 140px;"></select>
        </div>
        <div class="form-group col-md-2">
          <label for="ownership">Ownership</label>
          <select id="ownership" class="form-control">
            <option value="">Select</option>
            <option value="GOV">GOV (NMS)</option>
            <option value="PNFP">PNFP (JMS)</option>
          </select>
        </div>
        <div class="form-group col-md-2">
          <label for="uploadFinancialYear">Financial year</label>
          <input id="uploadFinancialYear" type="text" class="form-control" placeholder="e.g. 2025/26" autocomplete="off">
        </div>
        <div class="form-group col-md-3">
          <label for="xlsFile">File (XLS / XLSX)</label>
          <input id="xlsFile" type="file" class="form-control-file" accept=".xls,.xlsx">
        </div>
        <div class="form-group col-md-3 d-flex align-items-end">
          <button type="submit" class="btn btn-primary">Import</button>
        </div>
      </div>
    </form>
  </div>
</div>
{{end}}

{{define "extra_js"}}
<script>
(function() {
  var facilities = [];

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
    var $sel = $('#importRegion');
    $sel.find('option:not(:first)').remove();
    $sel.append($('<option value="" disabled>Loading…</option>'));
    apiGet('/regions', null, function(r) {
      $sel.find('option').remove();
      $sel.append($('<option value="">Select region</option>'));
      var list = (r && Array.isArray(r)) ? r : [];
      list.forEach(function(v) {
        var id = (v.id != null ? v.id : (v.ID != null ? v.ID : ''));
        var name = v.name || v.Name || '';
        if (name) $sel.append($('<option></option>').val(String(id)).attr('data-name', name).text(name));
      });
    });
  }

  function loadZones(regionId) {
    var $sel = $('#importZone');
    $sel.find('option:not(:first)').remove();
    $('#importDistrict').find('option:not(:first)').remove();
    $('#importFacility').find('option').remove();
    $('#importFacilitySearch').val('');
    regionId = regionId === undefined || regionId === null ? '' : String(regionId).trim();
    if (!regionId) {
      filterImportFacilities();
      return;
    }
    $sel.append($('<option value="" disabled>Loading…</option>'));
    apiGet('/zones', { region_id: regionId }, function(r) {
      $sel.find('option').remove();
      $sel.append($('<option value="">Select zone</option>'));
      var list = (r && Array.isArray(r)) ? r : [];
      list.forEach(function(v) {
        var id = (v.id != null ? v.id : (v.ID != null ? v.ID : ''));
        var name = v.name || v.Name || '';
        if (name) $sel.append($('<option></option>').val(String(id)).attr('data-name', name).text(name));
      });
      loadFacilitiesAfterLocation();
    });
  }

  function loadDistricts(zoneId, regionId) {
    var $sel = $('#importDistrict');
    $sel.find('option:not(:first)').remove();
    $('#importFacility').find('option').remove();
    $('#importFacilitySearch').val('');
    zoneId = zoneId === undefined || zoneId === null ? '' : String(zoneId).trim();
    regionId = regionId === undefined || regionId === null ? '' : String(regionId).trim();
    if (zoneId) {
      $sel.append($('<option value="" disabled>Loading…</option>'));
      apiGet('/districts', { zone_id: zoneId }, function(r) {
        $sel.find('option').remove();
        $sel.append($('<option value="">Select district</option>'));
        var list = (r && Array.isArray(r)) ? r : [];
        list.forEach(function(v) {
          var id = (v.id != null ? v.id : (v.ID != null ? v.ID : ''));
          var name = v.name || v.Name || '';
          if (name) $sel.append($('<option></option>').val(String(id)).attr('data-name', name).text(name));
        });
        loadFacilitiesAfterLocation();
      });
    } else if (regionId) {
      $sel.append($('<option value="" disabled>Loading…</option>'));
      apiGet('/districts', { region_id: regionId }, function(r) {
        $sel.find('option').remove();
        $sel.append($('<option value="">Select district</option>'));
        var list = (r && Array.isArray(r)) ? r : [];
        list.forEach(function(v) {
          var id = (v.id != null ? v.id : (v.ID != null ? v.ID : ''));
          var name = v.name || v.Name || '';
          if (name) $sel.append($('<option></option>').val(String(id)).attr('data-name', name).text(name));
        });
        loadFacilitiesAfterLocation();
      });
    } else {
      loadFacilitiesAfterLocation();
    }
  }

  function loadFacilitiesAfterLocation() {
    filterImportFacilities();
  }

  function loadLevelsOfCare() {
    apiGet('/levels-of-care', null, function(r) {
      var list = (r && Array.isArray(r)) ? r : [];
      var $sel = $('#importLevelOfCare');
      $sel.find('option').remove();
      $sel.append($('<option value="">Select level</option>'));
      list.forEach(function(v) {
        var id = (v.id != null ? v.id : (v.ID != null ? v.ID : ''));
        var name = v.name || v.Name || '';
        var code = v.code || v.Code || '';
        if (name) $sel.append($('<option></option>').val(String(id)).attr('data-name', name).attr('data-code', code).text(name));
      });
    });
  }

  function getRegionName() { return $('#importRegion option:selected').attr('data-name') || ''; }
  function getZoneName() { return $('#importZone option:selected').attr('data-name') || ''; }
  function getDistrictName() { return $('#importDistrict option:selected').attr('data-name') || ''; }

  function loadFacilities(regionName, districtName, levelCode) {
    var params = { active: 'true' };
    if (regionName) params.region = regionName;
    if (districtName) params.district = districtName;
    var $fac = $('#importFacility');
    $fac.find('option').remove();
    $fac.append($('<option value="" disabled>Loading…</option>'));
    apiGet('/facilities', params, function(r) {
      facilities = (r && Array.isArray(r)) ? r : [];
      filterImportFacilities();
    });
  }

  function getLevelCode() {
    return ($('#importLevelOfCare option:selected').attr('data-code') || '').trim();
  }
  function getLevelName() {
    return ($('#importLevelOfCare option:selected').attr('data-name') || '').trim();
  }

  function filterImportFacilities() {
    var regionName = (getRegionName() || '').trim();
    var districtName = (getDistrictName() || '').trim();
    var zoneName = (getZoneName() || '').trim();
    var levelCode = getLevelCode();
    var levelName = getLevelName();
    var q = ($('#importFacilitySearch').val() || '').toLowerCase().trim();
    var $sel = $('#importFacility');
    var selected = $sel.val();
    $sel.find('option').remove();
    var list = facilities.slice();
    if (regionName) {
      var rLower = regionName.toLowerCase();
      list = list.filter(function(f) {
        var fr = (f.region || f.Region || '').toString().trim().toLowerCase();
        return fr === rLower || fr.indexOf(rLower) >= 0 || rLower.indexOf(fr) >= 0;
      });
    }
    if (districtName) {
      var dLower = districtName.toLowerCase();
      list = list.filter(function(f) {
        var fd = (f.district || f.District || '').toString().trim().toLowerCase();
        return fd === dLower || fd.indexOf(dLower) >= 0 || dLower.indexOf(fd) >= 0;
      });
    }
    if (zoneName) {
      var zLower = zoneName.toLowerCase();
      list = list.filter(function(f) {
        var fz = (f.zone || f.Zone || '').toString().trim().toLowerCase();
        return fz === zLower || fz.indexOf(zLower) >= 0 || zLower.indexOf(fz) >= 0;
      });
    }
    if (levelCode || levelName) {
      var lc = (levelCode || '').toUpperCase().replace(/\s/g, '');
      var ln = (levelName || '').toLowerCase();
      list = list.filter(function(f) {
        var fl = (f.level_of_care || f.LevelOfCare || '').toString().trim();
        var flLower = fl.toLowerCase();
        var flCode = fl.toUpperCase().replace(/\s/g, '');
        return (lc && flCode === lc) || (ln && (flLower === ln || flLower.indexOf(ln) >= 0));
      });
    }
    if (list.length === 0 && facilities.length > 0) list = facilities.slice();
    list.filter(function(f) {
      var code = (f.facility_code || f.FacilityCode || '').toString().toLowerCase();
      var name = (f.facility_name || f.FacilityName || '').toString().toLowerCase();
      var label = code + ' - ' + name;
      return !q || code.indexOf(q) >= 0 || name.indexOf(q) >= 0 || label.indexOf(q) >= 0;
    }).forEach(function(f) {
      var code = (f.facility_code || f.FacilityCode || '');
      var name = (f.facility_name || f.FacilityName || '');
      var label = code + ' - ' + name;
      $sel.append($('<option></option>').val(code).attr('data-name', name).text(label));
    });
    if (selected && $sel.find('option[value="' + selected + '"]').length) $sel.val(selected);
  }

  $('#importRegion').on('change', function() {
    var id = $(this).val();
    if (!id) {
      $('#importZone').find('option:not(:first)').remove();
      $('#importDistrict').find('option:not(:first)').remove();
      loadZones('');
      loadDistricts(null, null);
      filterImportFacilities();
      return;
    }
    loadZones(id);
    loadDistricts(null, id);
    filterImportFacilities();
  });

  $('#importZone').on('change', function() {
    var zoneId = $(this).val();
    var regionId = $('#importRegion').val();
    if (!zoneId && !regionId) return;
    loadDistricts(zoneId || null, zoneId ? null : regionId);
    filterImportFacilities();
  });

  $('#importDistrict').on('change', function() {
    filterImportFacilities();
  });

  $('#importLevelOfCare').on('change', function() {
    filterImportFacilities();
  });

  $('#importFacilitySearch').on('input keyup', function() {
    filterImportFacilities();
  });

  $('#importFacility').on('change', function() {
    var opt = $(this).find('option:selected');
    if (opt.val()) $('#importFacilitySearch').val(opt.text());
  });

  $('#uploadProcurementXls').on('submit', function(e) {
    e.preventDefault();
    var fileInput = $('#xlsFile')[0];
    var facilityCode = ($('#importFacility').val() || '').toString().trim();
    var facilityName = $('#importFacility').find('option:selected').attr('data-name') || ($('#importFacilitySearch').val() || '').toString().trim();
    var fyInput = document.getElementById('uploadFinancialYear');
    var fy = (fyInput && fyInput.value) ? fyInput.value.trim() : '';
    var ownership = $('#ownership').val();
    if (!fileInput || !fileInput.files || !fileInput.files.length) {
      toastr.warning('Please choose an XLS or XLSX file.');
      return;
    }
    if (!facilityCode) {
      toastr.warning('Please select a facility from the list.');
      return;
    }
    if (!fy) {
      toastr.warning('Please enter the financial year (e.g. 2025/26).');
      return;
    }
    var fd = new FormData();
    fd.append('file', fileInput.files[0]);
    fd.append('facility_id', facilityCode);
    fd.append('facility_name', facilityName);
    fd.append('financial_year', fy);
    fd.append('region', getRegionName());
    fd.append('zone', getZoneName());
    fd.append('district', getDistrictName());
    var levelName = $('#importLevelOfCare option:selected').attr('data-name');
    if (levelName) fd.append('level_of_care', levelName);
    if (ownership) fd.append('ownership', ownership);
    $.ajax({
      url: '/api/v1/procurement-plans/upload-xls',
      method: 'POST',
      data: fd,
      contentType: false,
      processData: false,
      xhrFields: { withCredentials: true },
      success: function() {
        toastr.success('Import successful.');
        $('#xlsFile').val('');
      },
      error: function(xhr) {
        toastr.error(xhr.responseJSON && xhr.responseJSON.error ? xhr.responseJSON.error : 'Import failed.');
      }
    });
  });

  $(document).ready(function() {
    loadRegions();
    loadLevelsOfCare();
    loadFacilities('', '', '');
  });
})();
</script>
{{end}}
