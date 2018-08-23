$(function($) {
    $('#section-results-all').hide();
    $('#results-switch').change(function() {
        if (this.checked) {
            //show all results
            $('#section-results').hide();
            $('#section-results-all').show();
            
        } else {
            //show just normal results
            $('#section-results').show();
            $('#section-results-all').hide();
            
        }
    });
    
});