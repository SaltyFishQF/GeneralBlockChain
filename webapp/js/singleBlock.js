const hash = GetPar("hash")
var index

$(function () {
    $.ajax({
            url: `http://localhost:8080/block?hash=${hash}`,
            data: {},
            type: "get",
            success: function f(res) {
                initBlock(res)
                initTX()
            },
        }
    )
})

function GetPar(name) {
    let reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
    let r = window.location.search.substr(1).match(reg);
    if (r != null) return decodeURIComponent(r[2]);
    return null;
}

function initBlock(res) {
    let block = JSON.parse(res)
    index = block.Header.Index
    document.getElementById('bindex').innerHTML = `${block.Header.Index}`
    document.getElementById('bindex2').innerHTML = `${block.Header.Index}`
    document.getElementById('btime').innerHTML = `${block.Header.Timestamp}`
    document.getElementById('bhash').innerHTML = `${block.Hash}`
    document.getElementById('bprehash').innerHTML = `<a href="./../html/singleBlock.html?hash=${block.Header.PreviousHash}">${block.Header.PreviousHash}</a>`
    document.getElementById('btxnum').innerHTML = `${block.Header.Size}`
    document.getElementById('bmtroot').innerHTML = `${block.Header.MerkleRoot}`
}

function initTX() {
    $.ajax({
        url: `http://localhost:8080/transaction?id=${index}`,
        data: {},
        type: "get",
        success: function f(res) {
            let li = document.getElementById("txList")
            let list = JSON.parse(res)
            console.log(list)
            for (let i = 0; i < list.length; i++) {
                li.innerHTML += `<tr><td>${list[i].Address}</td><td>${list[i].From}</td><td>${list[i].To}</td><td>${list[i].InputData}</td></tr>`
            }
        },
    })
}