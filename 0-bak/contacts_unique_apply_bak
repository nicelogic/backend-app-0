
func (r *queryResolver) AddContactsApply(ctx context.Context, first *int, after *string) (*model.AddContactsApplyConnection, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("user: %#v query add contacts apply\n", user)

	variables := map[string]interface{}{
		"user_id": user.Id,
		"first":   first,
		"after":   after,
	}
	response, err := r.CassandraClient.Query(cassandra.AddContactsApplyGql, variables)
	if err != nil {
		return nil, err
	}
	fmt.Println(response)

	pageState, jsonValue, err := r.CassandraClient.QueryResponse(response)
	if err != nil {
		return nil, err
	}
	addContactsApplys := make([]model.AddContactsApply, 0)
	err = json.Unmarshal(jsonValue, &addContactsApplys)
	if err != nil {
		return nil, err
	}
	uniqueAddContactsApplys := make(map[string]*model.AddContactsApply)
	for _, apply := range addContactsApplys {
		apply := apply
		id := apply.UserID + ">" + apply.ContactsID
		uniqueApply := uniqueAddContactsApplys[id]
		if uniqueApply == nil || apply.UpdateTime > uniqueApply.UpdateTime {
			uniqueAddContactsApplys[id] = &apply
		}
	}

	addContactsApplyConnection := &model.AddContactsApplyConnection{}
	addContactsApplyConnection.TotalCount = len(uniqueAddContactsApplys)
	for _, apply := range uniqueAddContactsApplys {
		apply := apply
		edge := &model.AddContactsApplyEdge{}
		edge.Node = apply
		addContactsApplyConnection.Edges = append(addContactsApplyConnection.Edges, edge)
	}
	addContactsApplyConnection.PageInfo = &model.AddContactsApplyEdgePageInfo{}
	addContactsApplyConnection.PageInfo.EndCursor = pageState
	addContactsApplyConnection.PageInfo.HasNextPage = pageState != nil

	return addContactsApplyConnection, err
}
