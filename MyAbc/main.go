package main

import (
    "html/template"
    "log"
    //"os"
	"net/http"
)

type Page struct {
    Title []string
    Label []string
    Image []string
}

var templates = template.Must(template.ParseGlob("templates/*"))

func handler(w http.ResponseWriter, r *http.Request) {

    p := Page{
        Title: []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"},
        Label: []string{"Apple", "Bike", "Cat", "Dog", "Elephant", "Fish", "Goat", "Hen", "Ice-Cream", "Jug", "Kite", "Lion", "Monkey", "Nest", "Orange", "Peacock", "Queen", "Rat", "Swan", "Tiger", "Umbrela", "Van", "Window", "X-Mas tree", "Yatch", "Zebra"},
        Image: []string{"http://www.clker.com/cliparts/2/c/B/U/C/v/cartoon-red-apple.svg",
                        "https://i.ytimg.com/vi/LiBEDTt8hSE/hqdefault.jpg",
                        "https://i.ytimg.com/vi/Hw_iS1ouNZ4/hqdefault.jpg",
                        "https://rlv.zcache.com/cartoon_bernese_mountain_dog_cutout-r91d8e0a735c541da9b8627f042c42d3c_x7sai_8byvr_324.jpg",
                        "https://i.ytimg.com/vi/dSiLPiNVWTY/maxresdefault.jpg",
                        "https://thumb9.shutterstock.com/display_pic_with_logo/1058333/167223395/stock-vector-vector-illustration-of-angry-fish-cartoon-167223395.jpg",
                        "http://4vector.com/i/free-vector-cartoon-goat-clip-art_128387_cartoon-goat-clip-art/Cartoon_Goat_clip_art_hight.png",
                        "https://upload.wikimedia.org/wikipedia/commons/thumb/6/62/Hen_and_chicks_cartoon_04.svg/2000px-Hen_and_chicks_cartoon_04.svg.png",
                        "https://financialtribune.com/sites/default/files/field/image/17january/04-zs-ice-cream_66-ab.jpg",
                        "https://www.klevering.com/img/product/1190/0/0/jansen+co+my+jug+yellow.jpg",
                        "https://i.pinimg.com/originals/06/e0/64/06e0643ff49f576fa2ef2859e8e0a3f3.jpg",
                        "https://kids.nationalgeographic.com/content/dam/kids/photos/animals/Mammals/H-P/photoark-lion.ngsversion.1466004832449.adapt.945.1.png",
                        "https://www.longleat.co.uk/api/v2/storage/public/assets/25/monkey-mayhem/monkey-grid-4_original_original.jpg",
                        "https://e27.co/wp-content/uploads/2015/07/nest-blue-eggs.jpg",
                        "http://www.mulierchile.com/orange/orange-007.jpg",
                        "http://17rg073sukbm1lmjk9jrehb643.wpengine.netdna-cdn.com/wp-content/uploads/2014/05/800px-PeacockValencay.jpg",
                        "https://previews.123rf.com/images/tigatelu/tigatelu1509/tigatelu150900246/45091634-Cartoon-queen-holding-scepter-Stock-Vector.jpg",
                        "https://www.colourbox.com/preview/5192011-rat-cartoon.jpg",
                        "https://thumbs.dreamstime.com/z/cartoon-beauty-swan-floats-water-illustration-65446356.jpg",
                        "https://static6.depositphotos.com/1010340/558/v/950/depositphotos_5589085-stock-illustration-tiger-cartoon.jpg",
                        "https://guideimg.alibaba.com/images/shop/79/09/24/1/gifts-for-children-cartoon-frog-umbrella-sunshade-environmental-protection-umbrella-long-handle-umbrella-sun-rain-freeshipping_279441.jpg",
                        "https://thumbs.dreamstime.com/z/cartoon-retro-hippie-van-24431661.jpg",
                        "https://i.pinimg.com/736x/e0/e1/6c/e0e16cfd9886401603031b76eebf78b2--window-view-cartoon.jpg",
                        "https://thumbs.dreamstime.com/z/cartoon-christmas-tree-20742748.jpg",
                        "https://i.ytimg.com/vi/4v0BFgabYk4/maxresdefault.jpg",
                        "https://i.ytimg.com/vi/5xrfnong3wM/maxresdefault.jpg"},
}

    err := templates.ExecuteTemplate(w, "demo.html", p)
    if err != nil {
        log.Fatal("Cannot Get View ", err)
    }
}

func main() {

    http.HandleFunc("/view/", handler)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}