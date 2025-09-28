<?php

trait SQLiteDB
{
    private static $instance = null;

    public static function connectDB(?PDO $pdoInstance = null): PDO
    {
        if (is_null(static::$instance)) {
            static::$instance = $pdoInstance ?? new PDO('sqlite:./database/artbox.db');
        }

        return static::$instance;
    }
}

class DataRepository
{
    use SQLiteDB;

    /**
     * Get all oeuvres.
     */
    public static function getOeuvres(): PDOStatement
    {
        $db = static::connectDB();

        return $db->query('SELECT * FROM oeuvres');
    }

    /**
     * Get a single oeuvre by id.
     */
    public static function getOeuvreById(int $id): mixed
    {
        $db = static::connectDB();

        $stmt = $db->prepare('SELECT * FROM oeuvres WHERE id = :id');
        $stmt->execute([':id' => $id]);

        return $stmt->fetch(\PDO::FETCH_ASSOC);
    }

    /**
     * Insert a new oeuvre.
     */
    public static function insertOeuvre(array $fields): bool
    {
        $db = static::connectDB();

        $sql = <<<SQLREQUEST
        INSERT INTO oeuvres (titre, artiste, image, description)
        VALUES (:titre, :artiste, :image, :description)
        SQLREQUEST;

        $stmt = $db->prepare($sql);

        return $stmt->execute($fields);
    }
}
