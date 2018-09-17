var productSummary

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
            var product = data.taxable_product
            c = buildTaxableTableContent(product);
            $('#taxable-products-body').append(c);
            productSummary.total_tax_amount += product.tax_price;
            productSummary.total_amount += product.price;
            productSummary.grand_total += product.total_price;
            console.log(productSummary);
            loadTaxableSummary();
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

function loadTaxableSummary(){
    $("#total-amount").html("<b>"+productSummary.total_amount+"</b>");
    $("#total-tax-amount").html("<b>"+productSummary.total_tax_amount+"</b>");
    $("#grand-total").html("<b>"+productSummary.grand_total+"</b>");
}

function loadTaxableProducts(){
    $.ajax({
        type: "POST",
        url: "/shopee/tax/get_taxable_products",
        dataType: "json",
        success: function(data){
            productSummary = data.summary
            var contents = ""
            $.each(data.taxable_products, function(i, product){
                c = buildTaxableTableContent(product);
                contents += c;
            })
            $('#taxable-products-body').append(contents)
            loadTaxableSummary();
        }
    })
}