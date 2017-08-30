main = {

    /**
     * @param contents 显示内容
     * @param type info success warning danger rose primary
     * @param from  top center bottom
     * @param align left center right
     */
    notifyShow : function (contents, type, timer, from, align) {
        if(!from){
            from = 'top'
        }
        if(!align){
            align = 'right'
        }
        if(!timer){
            timer = 3000
        }

        if(!contents || !type){
            return
        }

        $.notify({
            icon: "notifications",
            message: contents

        },{
            type: type,
            timer: 3000,
            placement: {
                from: from,
                align: align
            }
        });
    },
}