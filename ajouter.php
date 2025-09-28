<?php 

require_once 'header.php'; 
require_once 'database.php';

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $isLengthValid = function (string $value, $len = 0) {
        return isset($value) && strlen(trim($value)) > $len;
    };
    $formFieldsValidator = [
        'titre' => [
            'errorMessage' => 'Le champ titre ne peut être vide.',
            'rule' =>  $isLengthValid
        ],
        'artiste' =>  [
            'errorMessage' => 'Le champ titre ne peut être vide.',
            'rule' =>  $isLengthValid
        ],
        'image' =>  [
            'errorMessage' => 'Le champ image doit être une URL.',
            'rule' =>  function (string $value) {
                return isset($value) && filter_var($value, FILTER_VALIDATE_URL);
            }
        ],
        'description' =>  [
            'errorMessage' => 'Le champ description doit contenir au moins 3 caractères.',
            'rule' => function (string $value) use ($isLengthValid) {
                return $isLengthValid($value, 3);
            }
        ],
    ];
    $errorMessages = [];
    $fields = [];

    foreach ($formFieldsValidator as $key => $validator) {
        if (!$validator['rule']($_POST[$key])) {
            $errorMessages[$key] = $validator['errorMessage'];
        }

        $fields[$key] = htmlentities($_POST[$key]);
    }

    if (empty($errorMessages) && DataRepository::insertOeuvre($fields)) {
        header('Location: index.php');
        exit();
    }
}
?>

<form action="ajouter.php" method="POST">
    <div class="champ-formulaire">
        <label for="titre">Titre de l'œuvre</label>
        <input type="text" name="titre" id="titre"
            value="<?= isset($fields['titre']) ? $fields['titre'] : '' ?>">
        <span class="error-message"><?= isset($errorMessages['titre']) ? $errorMessages['titre'] : '' ?></span>
    </div>
    <div class="champ-formulaire">
        <label for="artiste">Auteur de l'œuvre</label>
        <input type="text" name="artiste" id="artiste"
            value="<?= isset($fields['artiste']) ? $fields['artiste'] : '' ?>">
        <span class="error-message"><?= isset($errorMessages['artiste']) ? $errorMessages['artiste'] : '' ?></span>
    </div>
    <div class="champ-formulaire">
        <label for="image">URL de l'image</label>
        <input type="url" name="image" id="image"
            value="<?= isset($fields['image']) ? $fields['image'] : '' ?>">
        <span class="error-message"><?= isset($errorMessages['image']) ? $errorMessages['image'] : '' ?></span>
    </div>
    <div class="champ-formulaire">
        <label for="description">Description</label>
        <textarea name="description" id="description"><?=
            isset($fields['description']) ? $fields['description'] : ''
        ?></textarea>
        <span class="error-message"><?= isset($errorMessages['description']) ? $errorMessages['description'] : '' ?></span>
    </div>

    <input type="submit" value="Valider" name="submit">
</form>

<?php require_once 'footer.php'; ?>
