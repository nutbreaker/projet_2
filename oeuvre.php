<?php
    require_once 'header.php';
    require_once 'database.php';

    // Si l'URL ne contient pas d'id, on redirige sur la page d'accueil
    if(empty($_GET['id'])) {
        header('Location: index.php');
    }

    $db = DBConnection();
    $stmt = $db->prepare('SELECT * FROM oeuvres WHERE id = :id');
    $isSuccess = $stmt->execute([ ':id' => intval($_GET['id']) ]);

    if(!$isSuccess) {
        header('Location: index.php');
        exit();
    }

    $oeuvre = $stmt->fetch(\PDO::FETCH_ASSOC);

    // Si aucune oeuvre trouvé, on redirige vers la page d'accueil
    if(!$oeuvre) {
        header('Location: index.php');
        exit();
    }
?>

<article id="detail-oeuvre">
    <div id="img-oeuvre">
        <img src="<?= $oeuvre['image'] ?>" alt="<?= $oeuvre['titre'] ?>">
    </div>
    <div id="contenu-oeuvre">
        <h1><?= $oeuvre['titre'] ?></h1>
        <p class="description"><?= $oeuvre['artiste'] ?></p>
        <p class="description-complete">
             <?= $oeuvre['description'] ?>
        </p>
    </div>
</article>

<?php require_once 'footer.php'; ?>
