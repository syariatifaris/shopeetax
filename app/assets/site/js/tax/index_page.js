$(document).ready(function(){
    loadTaxableProducts();
});

function addTaxableProduct(){
    $.ajax({
        type: "POST",
        url: "/shopee/tax/insert",
        dataType: "json",
        contentType: "application/json",
        data: JSON.stringify({ 
            "tax_product": {
                "product_name": $("#add-product-name").val(), 
                "tax_category_id" : parseInt($("#add-tax-cat-id").val()), 
                "initial_price": parseFloat($("#add-tax-amount").val())
            }
        }), 
        success: function(data){
            c = buildTaxableTableContent(data.taxable_product);
            $('#taxable-products-body').append(c);
        },
        error: function(err){
            console.log(err);
        }
    })
}

function buildTaxableTableContent(product){
    var content = ""
    content += '<tr>'
    content += '<td>' +  product.product_name + '</td>';
    content += '<td>' +  product.tax_category_id + '</td>';
    content += '<td>' +  product.tax_code + '</td>';
    content += '<td>' +  product.price + '</td>';
    content += '<td>' +  product.tax_price + '</td>';
    content += '<td>' +  product.total_price + '</td>';
    content += '</tr>' 
    return content
}

function loadTaxableProducts(){
    $.ajax({
        type: "POST",
        url: "/shopee/tax/get_taxable_products",
        dataType: "json",
        success: function(data){
            var contents = ""
            $.each(data.taxable_products, function(i, product){
                c = buildTaxableTableContent(product);
                contents += c;
            })
            $('#taxable-products-body').append(contents)
        }
    })
}