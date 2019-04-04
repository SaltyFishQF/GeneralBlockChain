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
    index = block.header.index
    document.getElementById('bindex').innerHTML = `${block.header.index}`
    document.getElementById('btime').innerHTML = `${block.header.timestamp}`
    document.getElementById('bhash').innerHTML = `${block.hash}`
    document.getElementById('bprehash').innerHTML = `<a href="./../html/singleBlock.html?hash=${block.header.previousHash}">${block.header.previousHash}</a>`
    document.getElementById('btxnum').innerHTML = `${block.header.size}`
    document.getElementById('bmtroot').innerHTML = `${block.header.merkleRoot}`
}

function initTX() {
    $.ajax({
        url: `http://localhost:8080/transaction?id=${index}`,
        data: {},
        type: "get",
        success: function f(res) {
            let li = document.getElementById("txList")
            let list = JSON.parse(res)
            for (let i = 0; i < list.length; i++) {
                li.innerHTML += `<tr><td>${list[i].id}</td><td>${list[i].user}</td><td>${list[i].doc}</td><td>${list[i].inputData}</td></tr>`
            }
        },
    })
}