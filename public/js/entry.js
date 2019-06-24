var labours;
function onDateChange() {
  date = document.getElementById("date-picker");
  if (isDateProductPicked()) createEntryTable();
}

function onProductChange() {
  product = document.getElementById("product-picker");
  if (isDateProductPicked()) createEntryTable();
}

function isDateProductPicked() {
  product = document.getElementById("product-picker");
  date = document.getElementById("date-picker");
  return Boolean(date.value && product.value);
}

function onEntryFormSubmit() {
  var labour = document
    .getElementById("labour-name")
    .value.toUpperCase()
    .trim();
  var box = document.getElementById("box-no").value;
  var packet = document.getElementById("packet-no").value;
  var date = document.getElementById("date-picker").value;
  var product = document.getElementById("product-picker").value;
  var res = isDateProductPicked();
  var i = labours.indexOf(labour);
  if (i < 0) {
    if (!res) {
      alert("Pick Date and Product");
    } else {
      console.log("Entry to be Made");
      $.ajax({
        type: "POST",
        url:
          "/entry/new?labour=" +
          labour +
          "&box=" +
          box +
          "&packet=" +
          packet +
          "&date=" +
          date +
          "&product=" +
          product,
        success: function(s) {
          createEntryTable();
        },
        error: function() {
          console.log("faluire from server");
        }
      });
    }
  } else {
    alert("Name Already exist");
  }
  clearEntryForm();
}

function clearEntryForm() {
  nameControl = document.getElementById("labour-name");
  boxControl = document.getElementById("box-no");
  packetControl = document.getElementById("packet-no");
  nameControl.value = "";
  boxControl.value = "";
  packetControl.value = "";
  nameControl.focus();
}

function createEntryTable() {
  $("#journal-table").html("");
  var date = document.getElementById("date-picker").value;
  var product = document.getElementById("product-picker").value;
  $.ajax({
    type: "GET",
    url: "/entry/getall?date=" + date + "&id=" + product,
    success: function(p) {
      labours = [];
      let response = JSON.parse(p);
      entries = response.entries
      if (entries)
        entries.forEach(entry => {
          $("#journal-table").append(
            "<tr><td>" +
              entry.labour +
              "</td><td>" +
              entry.box +
              "</td>" +
              "<td>" +
              entry.packet +
              '</td><td><button onclick="RemoveEntry(' +
              entry.id +
              "," +
              entry.product +
              ')" ' +
              'class="btn btn-danger">Remove</button>'
          );
          labours = labours.concat(entry.labour);
        });
        totalmark = document.getElementById("total-entries")
        let box = response.tbox;
        let packet= response.tpacket
        if (box ==0 && packet == 0){
          totalmark.innerHTML = ""
        } else {
          totalmark.innerHTML = "Total: "+ box +" Boxes "+ packet+" Packets"
        }
    },
    error: function() {
      console.log("error in getiing jounral data");
    }
  });
}

function RemoveEntry(id, productID) {
  console.log(id, productID);
  $.ajax({
    type: "post",
    url: "/entry/delete?productID=" + productID + "&id=" + id,
    success: function(p) {
      console.log("product Deleted");
      createEntryTable();
    },
    error: function() {
      console.log("error in getiing jounral data");
    }
  });
}
