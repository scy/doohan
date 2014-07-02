(function ($) {
	"use strict";

	var entries = [];

	var $tpl      = $("#templates");
	var $entryTpl = $("#entryTpl");

	var $entries = $("#entries");

	var loadEntries = function () {
		$.getJSON("/1/entries", function (data) {
			entries = data;
			renderEntries();
		});
	};

	var renderEntry = function (entry) {
		var $entry = $entryTpl.clone();
		$entry.attr("id", null);
		$entry.find('[data-key="description"]').text(entry.description);
		return $entry;
	};

	var renderEntries = function () {
		$entries.empty();
		$.each(entries, function (idx, entry) {
			renderEntry(entry).appendTo($entries);
		});
	};

	loadEntries();

	$("form").submit(function () {
		var $desc = $('[name="description"]');
		$.post("/1/entries", { description: $desc.val() }, function (data) {
			loadEntries();
			$desc.val("");
		});
		return false;
	});
})(jQuery);
