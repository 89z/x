package youtube

type Player struct{
   Description struct{
      SimpleText string
   }
   PublishDate string
   Title struct{
      SimpleText string
   }
   ViewCount string
}

type response struct{
   Microformat struct{
      PlayerMicroformatRenderer Player
   }
}
