package internal

type Image struct {
	Src         string
	Description string
}

type Blog struct {
	Title  string
	Images []*Image
}

var Blogs = []*Blog{
	{
		Title: "We had a date at the MAP",
		Images: []*Image{
			{
				Src:         "https://vincent-1316174341.cos.ap-nanjing.myqcloud.com/website/us.jpeg",
				Description: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
			},
			{
				Src: "https://vincent-1316174341.cos.ap-nanjing.myqcloud.com/website/Statue%20of%20Venus%20Callipyge.jpeg",
				//Description: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
			},
		},
	},
}
