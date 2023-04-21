package seeders
import (

    "github.com/gofiber/fiber/v2"
    "github.com/diegocheca/goPosts.git/models"
    "github.com/diegocheca/goPosts.git/database"
    
)


var slugsList []string = []string{ "Aaren" , "Aarika" , "Abagael" , "Abagail" , "Abbe" , "Abbey" , "Ronny" , "Roobbie" , "Rora" , "Rori" , "Rorie" , "Rory" , "Ros" , "Rosa" , "Rosabel" , "Rosabella" , "Rosabelle" , "Rosaleen" , "Rosalia" , "Rosalie" , "Rosalind" , "Rosalinda" , "Rosalinde" , "Rosaline" , "Rosalyn" , "Rosalynd" , "Rosamond" , "Rosamund" , "Rosana" , "Rosanna" , "Rosanne" , "Rose" , "Roseann" , "Roseanna" , "Roseanne" , "Roselia" , "Roselin" , "Roseline" , "Rosella" , "Roselle" , "Rosemaria" , "Rosemarie" , "Rosemary" , "Rosemonde" , "Rosene" , "Rosetta" , "Rosette" , "Roshelle" , "Rosie" , "Rosina" , "Rosita" , "Roslyn" , "Rosmunda" , "Rosy" , "Row" , "Rowe" , "Rowena" , "Roxana" , "Roxane" , "Roxanna" , "Roxanne" , "Roxi" , "Roxie" , "Roxine" , "Roxy" , "Roz" , "Rozalie" , "Rozalin" , "Rozamond" , "Rozanna" , "Rozanne" , "Roze" , "Rozele" , "Rozella" , "Rozelle" , "Rozina" , "Rubetta" , "Rubi" , "Rubia" , "Rubie" , "Rubina" , "Ruby" , "Ruperta" , "Ruth" , "Ruthann" , "Ruthanne" , "Ruthe" , "Ruthi" , "Ruthie" , "Ruthy" , "Ryann" , "Rycca" , "Saba" , "Sabina" , "Sabine" , "Sabra" , "Sabrina" , "Sacha" , "Sada" , "Sadella" , "Sadie" , "Sadye" , "Saidee" , "Sal" , "Salaidh" , "Sallee" , "Salli" , "Sallie" , "Sally" , "Sallyann" , "Sallyanne" , "Saloma" , "Salome" , "Salomi" , "Sam" , "Samantha" , "Samara" , "Samaria" , "Sammy" , "Sande" , "Sandi" , "Sandie" , "Sandra" , "Sandy" , "Sandye" , "Sapphira" , "Sapphire" , "Sara" , "Sara-Ann" , "Saraann" , "Sarah" , "Sarajane" , "Saree" , "Sarena" , "Sarene" , "Sarette" , "Sari" , "Sarina" , "Sarine" , "Sarita" , "Sascha" , "Sasha" , "Sashenka" , "Saudra" , "Saundra" , "Savina" , "Sayre" , "Scarlet" , "Scarlett" , "Sean" , "Seana" , "Seka" , "Sela" , "Selena" , "Selene" , "Selestina" , "Selia" , "Selie" , "Selina" , "Selinda" , "Seline" , "Sella" , "Selle" , "Selma" , "Sena" , "Sephira" , "Serena" , "Serene" , "Shae" , "Shaina" , "Shaine" , "Shalna" , "Shalne" , "Shana" , "Shanda" , "Shandee" , "Shandeigh" , "Shandie" , "Shandra" , "Shandy" , "Shane" , "Shani" , "Shanie" , "Shanna" , "Shannah" , "Shannen" , "Shannon" , "Shanon" , "Shanta" , "Shantee" , "Shara" , "Sharai" , "Shari" , "Sharia" , "Sharity" , "Sharl" , "Sharla" , "Sharleen" , "Sharlene" , "Sharline" , "Sharon" , "Sharona" , "Sharron" , "Sharyl" , "Shaun" , "Shauna" , "Shawn" , "Shawna" , "Shawnee" , "Shay" , "Shayla" , "Shaylah" , "Shaylyn" , "Shaylynn" , "Shayna" , "Shayne" , "Shea" , "Sheba" , "Sheela" , "Sheelagh" , "Sheelah" , "Sheena" , "Sheeree" , "Sheila" , "Sheila-Kathryn" , "Sheilah" , "Shel" , "Shela" , "Shelagh" , "Shelba" , "Shelbi" , "Shelby" , "Shelia" , "Shell" , "Shelley" , "Shelli" , "Shellie" , "Shelly" , "Shena" , "Sher" , "Sheree" , "Sheri" , "Sherie" , "Sherill" , "Sherilyn" , "Sherline" , "Sherri" , "Sherrie" , "Sherry" , "Sherye" , "Sheryl" , "Shina" , "Shir" , "Shirl" , "Shirlee" , "Shirleen" , "Shirlene" , "Shirley" , "Shirline" , "Shoshana" , "Shoshanna" , "Siana" , "Sianna" , "Sib" , "Sibbie" , "Sibby" , "Sibeal" , "Sibel" , "Sibella" , "Sibelle" , "Sibilla" , "Sibley" , "Sibyl" , "Sibylla" , "Sibylle" , "Sidoney" , "Sidonia" , "Sidonnie" , "Sigrid" , "Sile" , "Sileas" , "Silva" , "Silvana" , "Silvia" , "Silvie" , "Simona" , "Simone" , "Simonette" , "Simonne" , "Sindee" , "Siobhan" , "Sioux" , "Siouxie" , "Sisely" , "Sisile" , "Sissie" , "Sissy" , "Siusan" , "Sofia" , "Sofie" , "Sondra" , "Sonia" , "Sonja" , "Sonni" , "Sonnie" , "Sonnnie" , "Sonny" , "Sonya" , "Sophey" , "Sophi" , "Sophia" , "Sophie" , "Sophronia" , "Sorcha" , "Sosanna" , "Stace" , "Stacee" , "Stacey" , "Staci" , "Stacia" , "Stacie" , "Stacy" , "Stafani" , "Star" , "Starla" , "Starlene" , "Starlin" , "Starr" , "Stefa" , "Stefania" , "Stefanie" , "Steffane" , "Steffi" , "Steffie" , "Stella" , "Stepha" , "Stephana" , "Stephani" , "Stephanie" , "Stephannie" , "Stephenie" , "Stephi" , "Stephie" , "Stephine" , "Stesha" , "Stevana" , "Stevena" , "Stoddard" , "Storm" , "Stormi" , "Stormie" , "Stormy" , "Sue" , "Suellen" , "Sukey" , "Suki" , "Sula" , "Sunny" , "Sunshine" , "Susan" , "Susana" , "Susanetta" , "Susann" , "Susanna" , "Susannah" , "Susanne" , "Susette" , "Susi" , "Susie" , "Susy" , "Suzann" , "Suzanna" , "Suzanne" , "Suzette" , "Suzi" , "Suzie" , "Suzy" , "Sybil" , "Sybila" , "Sybilla" , "Sybille" , "Sybyl" , "Sydel" , "Sydelle" , "Sydney" , "Sylvia" , "Tabatha" , "Tabbatha" , "Tabbi" , "Tabbie" , "Tabbitha" , "Tabby" , "Tabina" , "Tabitha" , "Taffy" , "Talia" , "Tallia" , "Tallie" , "Tallou" , "Tallulah" , "Tally" , "Talya" , "Talyah" , "Tamar" , "Tamara" , "Tamarah" , "Tamarra" , "Tamera" , "Tami" , "Tamiko" , "Tamma" , "Tammara" , "Tammi" , "Tammie" , "Tammy" , "Tamqrah" , "Tamra" , "Tana" , "Tandi" , "Tandie" , "Tandy" , "Tanhya" , "Tani" , "Tania" , "Tanitansy" , "Tansy" , "Tanya" , "Tara" , "Tarah" , "Tarra" , "Tarrah" , "Taryn" , "Tasha" , "Tasia" , "Tate" , "Tatiana" , "Tatiania" , "Tatum" , "Tawnya" , "Tawsha" , "Ted" , "Tedda" , "Teddi" , "Teddie" , "Teddy" , "Tedi" , "Tedra" , "Teena" , "TEirtza" , "Teodora" , "Tera" , "Teresa" , "Terese" , "Teresina" , "Teresita" , "Teressa" , "Teri" , "Teriann" , "Terra" , "Terri" , "Terrie" , "Terrijo" , "Terry" , "Terrye" , "Tersina" , "Terza" , "Tess" , "Tessa" , "Tessi" , "Tessie" , "Tessy" , "Thalia" , "Thea" , "Theadora" , "Theda" , "Thekla" , "Thelma" , "Theo" , "Theodora" , "Theodosia" , "Theresa" , "Therese" , "Theresina" , "Theresita" , "Theressa" , "Therine" , "Thia" , "Thomasa" , "Thomasin" , "Thomasina" , "Thomasine" , "Tiena" , "Tierney" , "Tiertza" , "Tiff" , "Tiffani" , "Tiffanie" , "Tiffany" , "Tiffi" , "Tiffie" , "Tiffy" , "Tilda" , "Tildi" , "Tildie" , "Tildy" , "Tillie" , "Tilly" , "Tim" , "Timi" , "Timmi" , "Timmie" , "Timmy" , "Timothea" , "Tina" , "Tine" , "Tiphani" , "Tiphanie" , "Tiphany" , "Tish" , "Tisha" , "Tobe" , "Tobey" , "Tobi" , "Toby" , "Tobye" , "Toinette" , "Toma" , "Tomasina" , "Tomasine" , "Tomi" , "Tommi" , "Tommie" , "Tommy" , "Toni" , "Tonia" , "Tonie" , "Tony" , "Tonya" , "Tonye" , "Tootsie" , "Torey" , "Tori" , "Torie" , "Torrie" , "Tory" , "Tova" , "Tove" , "Tracee" , "Tracey" , "Traci" , "Tracie" , "Tracy" , "Trenna" , "Tresa" , "Trescha" , "Tressa" , "Tricia" , "Trina" , "Trish" , "Trisha" , "Trista" , "Trix" , "Trixi" , "Trixie" , "Trixy" , "Truda" , "Trude" , "Trudey" , "Trudi" , "Trudie" , "Trudy" , "Trula" , "Tuesday" , "Twila" , "Twyla" , "Tybi" , "Tybie" , "Tyne" , "Ula" , "Ulla" , "Ulrica" , "Ulrika" , "Ulrikaumeko" , "Ulrike" , "Umeko" , "Una" , "Ursa" , "Ursala" , "Ursola" , "Ursula" , "Ursulina" , "Ursuline" , "Uta" , "Val" , "Valaree" , "Valaria" , "Vale" , "Valeda" , "Valencia" , "Valene" , "Valenka" , "Valentia" , "Valentina" , "Valentine" , "Valera" , "Valeria" , "Valerie" , "Valery" , "Valerye" , "Valida" , "Valina" , "Valli" , "Vallie" , "Vally" , "Valma" , "Valry" , "Van" , "Vanda" , "Vanessa" , "Vania" , "Vanna" , "Vanni" , "Vannie" , "Vanny" , "Vanya" , "Veda" , "Velma" , "Velvet" , "Venita" , "Venus" , "Vera" , "Veradis" , "Vere" , "Verena" , "Verene" , "Veriee" , "Verile" , "Verina" , "Verine" , "Verla" , "Verna" , "Vernice" , "Veronica" , "Veronika" , "Veronike" , "Veronique" , "Vevay" , "Vi" , "Vicki" , "Vickie" , "Vicky" , "Victoria" , "Vida" , "Viki" , "Vikki" , "Vikky" , "Vilhelmina" , "Vilma" , "Vin" , "Vina" , "Vinita" , "Vinni" , "Vinnie" , "Vinny" , "Viola" , "Violante" , "Viole" , "Violet" , "Violetta" , "Violette" , "Virgie" , "Virgina" , "Virginia" , "Virginie" , "Vita" , "Vitia" , "Vitoria" , "Vittoria" , "Viv" , "Viva" , "Vivi" , "Vivia" , "Vivian" , "Viviana" , "Vivianna" , "Vivianne" , "Vivie" , "Vivien" , "Viviene" , "Vivienne" , "Viviyan" , "Vivyan" , "Vivyanne" , "Vonni" , "Vonnie" , "Vonny" , "Vyky" , "Wallie" , "Wallis" , "Walliw" , "Wally" , "Waly" , "Wanda" , "Wandie" , "Wandis" , "Waneta" , "Wanids" , "Wenda" , "Wendeline" , "Wendi" , "Wendie" , "Wendy" , "Wendye" , "Wenona" , "Wenonah" , "Whitney" , "Wileen" , "Wilhelmina" , "Wilhelmine" , "Wilie" , "Willa" , "Willabella" , "Willamina" , "Willetta" , "Willette" , "Willi" , "Willie" , "Willow" , "Willy" , "Willyt" , "Wilma" , "Wilmette" , "Wilona" , "Wilone" , "Wilow" , "Windy" , "Wini" , "Winifred" , "Winna" , "Winnah" , "Winne" , "Winni" , "Winnie" , "Winnifred" , "Winny" , "Winona" , "Winonah" , "Wren" , "Wrennie" , "Wylma" , "Wynn" , "Wynne" , "Wynnie" , "Wynny" , "Xaviera" , "Xena" , "Xenia" , "Xylia" , "Xylina" , "Yalonda" , "Yasmeen" , "Yasmin" , "Yelena" , "Yetta" , "Yettie" , "Yetty" , "Yevette" , "Ynes" , "Ynez" , "Yoko" , "Yolanda" , "Yolande" , "Yolane" , "Yolanthe" , "Yoshi" , "Yoshiko" , "Yovonnda" , "Ysabel" , "Yvette" , "Yvonne" , "Zabrina" , "Zahara" , "Zandra" , "Zaneta" , "Zara" , "Zarah" , "Zaria" , "Zarla" , "Zea" , "Zelda" , "Zelma" , "Zena" , "Zenia" , "Zia" , "Zilvia" , "Zita" , "Zitella" , "Zoe" , "Zola" , "Zonda" , "Zondra" , "Zonnya" , "Zora" , "Zorah" , "Zorana" , "Zorina" , "Zorine" , "Zsa Zsa" , "Zsazsa" , "Zulema" , "Zuzana" }
var contentList []string = []string{"You can now derive a schema from example CSV, JSON, or XML data." ,"You can now generate your own custom data types using AI." ,"You can now generate fields on any topic using AI" ,"Added support for XML attributes by naming fields starting with 2" ,"Added the ability to generate v5 UUIDs via a new uuid_v5(namespace, name) function in formulas" ,"Added the ability to force the quote character on custom file formats." ,"Added Address Line 2 type." ,"Added a ULID data type." ,"Added airport data." ,"You can now generate datasets using JSON and import them into other schemas using the Dataset Column type." ,"Added support for InfluxDB" ,"Added the ability to import fields from a JSON schema or example JSON object." ,"You can now create a dataset directly from a schema. You no longer need to download and reupload generated data to create a dataset!" ,"You can now stream data to an MQTT endpoint! Click More > Stream to MQTT Endpoint... to get started." ,"Added types related to construction work." ,"Added Etherium and Tezos types."}
func randomSlug() (res string) {
	return randomPickStr(slugsList)
}


func randomContent() (res string) {
	return randomPickStr(contentList)
}

// randomSample picks a random element from array arr
func randomPickStr(arr []string) string {

	return arr[rand.Intn(len(arr))]
}




func getValidPost() (post Post) {
	post := new(models.Post)
	post.Image = "https://source.unsplash.com/random/?city,night"
	post.Thumbnail = "https://source.unsplash.com/random/?city,night"
	post.Slug = randomSlug()
	post.Title = randomSlug()
	post.Subtitle = randomSlug()
    post.Content = randomContent()
	post.Author = rand.Intn()
	post.Rate = rand.Intn()
	return post
}




func main() {

	NDATA := 8000
	for i := 0; i < NDATA; i++ {
		post := getValidPost()

		result := database.DB.Db.Create(&post)

		if i%100 == 99 {
			fmt.Printf("Seeded %d data\n", i+1)
		}
	}

	fmt.Printf("Successfully added %d data to database\n", NDATA)
}

