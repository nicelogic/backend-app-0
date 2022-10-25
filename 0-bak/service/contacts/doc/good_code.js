
export const getPaginatedItems = async (options: PaginationOptions) => {
	// options for page of pagination
	// it is encoded
     const { first, after } = options;
     const criteria = after
	 ? {
		 _id: {
		     $lt: decode(options.after),
		 },
	   }
	 : {};
     let items: Item[] = await mongoose
	 .model('item')
	 .find(criteria)
	 .sort({ _id: -1 })
	 .limit(options.first + 1)
	 .lean()
	 .exec();
     const hasNextPage = items.length > first - 1;
     //remove extra
     if (hasNextPage) {
	 items= items.slice(0, items.length - 1);
     }
     const edges = items.map(r => ({
	 cursor: encode(r._id.toString()),
	 node: r,
     }));
     return {
	 pageInfo: {
	     hasNextPage,
	 },
	 edges,
     };
 };