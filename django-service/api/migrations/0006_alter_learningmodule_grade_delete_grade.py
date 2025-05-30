# Generated by Django 5.2.1 on 2025-05-11 12:06

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("api", "0005_grade_alter_learningmodule_grade"),
    ]

    operations = [
        migrations.AlterField(
            model_name="learningmodule",
            name="grade",
            field=models.CharField(
                choices=[
                    ("A*", "A*"),
                    ("A", "A"),
                    ("B+", "B+"),
                    ("B", "B"),
                    ("C+", "C+"),
                    ("C", "C"),
                    ("D+", "D+"),
                    ("D", "D"),
                    ("E", "E"),
                    ("F", "F"),
                    ("S", "S"),
                    ("X", "X"),
                ],
                max_length=2,
            ),
        ),
        migrations.DeleteModel(
            name="Grade",
        ),
    ]
