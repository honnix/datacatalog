package filters

const (
	partitionJoinQuery string = "JOIN artifacts.artifact_id on partitions.artifact_id"
	tagsJoinQuery      string = "JOIN artifacts.artifact_id on tags.artifact_id"
)

func construct_query() string {

}


interface GormFilterExpression {
	GetEntity()
}
