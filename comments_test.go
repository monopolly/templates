package templates

//testing

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveComments(t *testing.T) {
	function, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(function).Name()
	fmt.Println(fmt.Sprintf("%s\n", fn[strings.LastIndex(fn, ".Test")+5:]))
	a := assert.New(t)
	_ = a

	html := []byte(`
	<div>




				<!-- comment1 -->
        <!-- comment2 -->
    
        <a>OK HTML</a>
        <!-- multicomment
        multicomment
        multicomment
        multicomment
        multicomment
				multicomment -->

				<script>
				//js comment
				/* js multicomment
				multicomment
				multicomment
				multicomment
				multicomment
				multicomment */
				OK js //comment
				</script>

				</div>
				</div>
				</div>

		<!-- <div class="center">
				<p class="desc">тут все скроллится)</p>
		</div> -->
<!-- filters -->
<div class="center" style="padding-top: 50px;">
<!-- <div class="filters"></div> -->


<!-- <div class="filters"></div> -->

<!-- 
<div style="overflow: hidden; width: 100vw;">
		<div class="row removeScrollBar scroll">
				<p :class="{active: query.category == key}" v-for="key in dict.jobs.categorySort" :key="key" class="tag" @click="query.category = key; popup_job_category = false; search()"> {{dict.jobs.category[key]}}</p> </div>
		<div class="row removeScrollBar scroll">
				<p v-for="(item, index) in dict.jobs.benefits" :key="index" class="tag" :class="{active: exist(query.benefits, index)}" @mousedown="push(query.benefits,index); search()">{{item}}</p>
		</div>

		<div class="row removeScrollBar scroll">
						<p class="tag" :class="{active : query.city == undefined && query.place == 0}" @click="query.place == 0 ? query.place = undefined : query.city = undefined; query.place=0; popup_job_city = false; search()">Удаленно</p>
						<p class="tag" :class="{active:query.city == index && query.place > 0}" v-for="(item, index) in dict.city" :key="index" @click="query.city == index ? query.city = undefined : query.city = index; query.place = 1; popup_job_city = false; search()">{{item}}</p>        
				</div>                
</div>

<br><br><br><br><br> -->
<div class="searchFilters">



			//moscow
        [
        {name:'Экспоцентр',     metro:4, station:11, adress : 'Краснопресненская наб., 14',          link:'http://expocentr.ru/'},
        {name:'Крокус Экспо',   metro:3, station:18, adress : 'ул. Международная, 16',                     link:'http://crocus-expo.ru/'},
        {name:'СК Олимпийский', metro:5, station:7, adress : 'пр. Олимпийский, 16',                       link:'http://olimpik.ru'},
        {name:'Artplay',        metro:0, station:5, adress : 'ул. Нижняя Сыромятническая, 10',            link:'http://artplay.ru'},
        {name:'ВДНХ',           metro:5, station:4, adress : 'пр. Мира, 119',                        link:'http://expo.vdnh.ru'},
        {name:'Гостиный двор',  metro:5, station:10, adress : 'ул. Ильинка, 4',                            link:'http://mosgd.ru'},
        {name:'ЦДХ',            metro:0, station:0, adress : 'ул. Крымский вал, 10',                          link:'http://cha.ru'},
        {name:'Deworkacy',      metro:1, station:11, adress : 'ул. Берсеньевская наб. 6с3',                    link:'http://deworkacy.ru'},
        {name:'Манеж',          metro:1, station:10, adress : 'ул. Манежная пл., 1',                           link:'http://expo.vdnh.ru'},
        {name:'Сокольники',     metro:1, station:3, adress : '5-й Лучевой просек, дом 7, строение 1',     link:'http://sokolniki.com/'},
        {name:'ЦМТ',            metro:6, station:8, adress : 'ул. Краснопресненская наб. 12, подъезд №4',     link:'http://wtcmoscow.ru/'},
        {name:'Сколково',       metro:4, station:2, adress : 'ул. Краснопресненская наб. 12, подъезд №4',     link:'http://sk.ru/'},
        ],
        //piter


			

</div>`)

	html = removeHTMLComments(html)
	html = removeJSComments(html)
	html = removeSingleJSComments(html)
	//fmt.Println(string(html))

}
