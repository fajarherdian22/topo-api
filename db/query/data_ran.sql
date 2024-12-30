-- name: GetAllByLevel :many
SELECT * FROM bh_data_level
WHERE `level` = ?
ORDER BY level_name asc;

-- name: GetByLevelAndName :many
SELECT * FROM bh_data_level
WHERE `level` = ? AND level_name = ?
ORDER BY level_name asc;

-- name: GetByLevelAndReference :many
SELECT * FROM bh_data_level
WHERE `level` = ? AND reference_name = ?
ORDER BY level_name asc;

-- name: ListDistinctNamesByLevel :many
SELECT DISTINCT(level_name) FROM bh_data_level
WHERE `level` = ?
ORDER BY level_name asc;
