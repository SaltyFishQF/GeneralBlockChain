const hash = GetPar("hash")

$(function () {
    $.ajax({
            url: `http://localhost:8080/block?hash=${hash}`,
            data: {},
            type: "get",
            success: function f(res) {
                initBlock(res)
            },
        }
    )
})

function GetPar(name) {
    let reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
    let r = window.location.search.substr(1).match(reg);
    if(r != null) return decodeURIComponent(r[2]);
    return null;
}

function initBlock(res){
    let block = JSON.parse(res)
    document.getElementById('bindex').innerHTML = `${block.header.index}`
    document.getElementById('btime').innerHTML = `${block.header.timestamp}`
    document.getElementById('bhash').innerHTML = `${block.hash}`
    document.getElementById('bprehash').innerHTML = `${block.header.previousHash}`
    document.getElementById('btxnum').innerHTML = `${block.header.size}`
    document.getElementById('bmtroot').innerHTML = `${block.header.merkleRoot}`
}