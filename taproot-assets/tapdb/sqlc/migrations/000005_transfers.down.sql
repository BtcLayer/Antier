DROP INDEX IF EXISTS transfer_time_idx;
DROP INDEX IF EXISTS transfer_txn_idx;
DROP TABLE IF EXISTS asset_transfers;
DROP TABLE IF EXISTS asset_deltas;
DROP TABLE IF EXISTS transfer_proofs;
DROP INDEX IF EXISTS transfer_inputs_idx;
DROP TABLE IF EXISTS asset_transfer_inputs;
DROP INDEX IF EXISTS transfer_outputs_idx;
DROP TABLE IF EXISTS asset_transfer_outputs;
DROP INDEX IF EXISTS proof_locator_hash_index;
DROP TABLE IF EXISTS receiver_proof_transfer_attempts;
DROP INDEX IF EXISTS passive_assets_idx;
DROP TABLE IF EXISTS passive_assets;
