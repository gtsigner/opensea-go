package opensea

import (
	"context"

	"github.com/pinealctx/opensea-go/model"
	"github.com/pinealctx/restgo"
)

// Assets To retrieve assets from our API, call the /assets endpoint with the desired filter parameters.
func (c *Client) Assets(ctx context.Context, req *AssetsRequest) (*AssetsResponse, error) {
	var rsp, err = c.get(ctx, "/api/v1/assets", restgo.ObjectParams(req)...)
	if err != nil {
		return nil, err
	}
	var response AssetsResponse
	err = ParseRsp(rsp, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type AssetsRequest struct {
	// The address of the owner of the assets
	Owner string `query:"owner"`
	// An array of token IDs to search for (e.g. ?token_ids=1&token_ids=209).
	// Will return a list of assets with token_id matching any of the IDs in this array.
	TokenIDs []string `query:"token_ids"`
	// Limit responses to members of a collection.
	// Case-sensitive and must match the collection slug exactly.
	// Will return all assets from all contracts in a collection.
	// For more information on collections, see our collections documentation.
	Collection string `query:"collection"`
	// Limit responses to members of a collection.
	// Case sensitive and must match the collection slug exactly.
	// Will return all assets from all contracts in a collection.
	// For more information on collections,
	CollectionSlug string `query:"collection_slug"`
	// Collection editor
	CollectionEditor string `query:"collection_editor"`
	// Can be asc for ascending or desc for descending
	OrderDirection string `query:"order_direction,required"`
	// The NFT contract address for the assets
	AssetContractAddress string `query:"asset_contract_address"`
	// An array of contract addresses to search for (e.g. ?asset_contract_addresses=0x1...&asset_contract_addresses=0x2...).
	// Will return a list of assets with contracts matching any of the addresses in this array.
	// If "token_ids" is also specified, then it will only return assets that match each (address, token_id) pairing, respecting order.
	AssetContractAddresses []string `query:"asset_contract_addresses"`
	// Limit
	Limit int32 `query:"limit,required"`
	//Cursor A cursor pointing to the page to retrieve
	Cursor string `query:"cursor"`
	// A flag determining if order information should be included in the response.
	IncludeOrders bool `query:"include_orders"`
}

type AssetsResponse struct {
	// List of Asset Object
	Assets []*model.Asset `opensea:"assets"`
	// A cursor to be supplied as a query param to retrieve the next page
	Next string `opensea:"next"`
	// A cursor to be supplied as a query param to retrieve the previous page
	Previous string `opensea:"previous"`
}
