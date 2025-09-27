<?php
if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $isLengthValid = function (string $value, $len = 0) {
        return isset($value) && strlen($value) > $len;
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
}

require_once 'ajouter.php';
