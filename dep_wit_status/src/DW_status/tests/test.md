func init() {
   // 需要在init中注册定义的model
   orm.RegisterModel(new(TwitterRelationship))
}

type TwitterRelationship struct {
    Id            string `orm:"column(uid);pk"` // 设置主键
    User          string
    RepostUser    string
    AtUser        string
    RepostLevel   string
    AtLevel       string
    TwitterRelationships    *TwitterRelationship `orm:"rel(fk)"`
}




func init() {
      // 需要在init中注册定义的model
      orm.RegisterModel(new(Relationship))
}


type Relationship struct {
      Id            string `orm:"column(uid);pk"` // 设置主键
      User          string
      RepostUser    string
      AtUser        string
      RepostLevel   string
      AtLevel       string
      Relationships    *Relationship `orm:"rel(fk)"`
}
