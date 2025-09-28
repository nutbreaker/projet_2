<?php require_once 'header.php'; ?>

<form action="traitement.php" method="POST">
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
